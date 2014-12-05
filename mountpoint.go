package main

import (
	"github.com/dustin/go-humanize"
	"syscall"
)

type MountPoint struct {
	percentAvail uint8
	freeSpace    string // Output from humanize.
	totalSize    string // Output frin humanize.
	mountPoint   string
}

// MountPointData returns statfs information for each mount point in a channel.
func MountPointData() <-chan *MountPoint {
	out := make(chan *MountPoint)
	go func() {
		for _, mountPoint := range Config.Check.Mountpoint {
			statfs, err := MountPointStatus(mountPoint)
			if err != nil {
				Logger.Printf("Unable to get statfs data for mount point %q: %v", mountPoint, err)
			}
			out <- &MountPoint{
				percentAvail: PercentAvailable(statfs),
				freeSpace:    humanize.Bytes(statfs.Bavail * uint64(statfs.Bsize)),
				totalSize:    humanize.Bytes(statfs.Blocks * uint64(statfs.Bsize)),
				mountPoint:   mountPoint,
			}
		}
		close(out)
	}()
	return out
}

// MountPointStatus checks specified mount point and returns its status.
func MountPointStatus(mountpoint string) (*syscall.Statfs_t, error) {
	var statfs syscall.Statfs_t
	err := syscall.Statfs(mountpoint, &statfs)
	return &statfs, err
}

// PercentAvailable calculates how many percent are available given input statfs data.
func PercentAvailable(statfs *syscall.Statfs_t) uint8 {
	if statfs.Blocks == 0 {
		return 0
	}
	return uint8(float32(statfs.Bavail) / float32(statfs.Blocks) * 100.0)
}
