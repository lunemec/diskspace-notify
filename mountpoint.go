package main

import (
	"syscall"
)

type MountPointData struct {
	percentAvail float32
	mountPoint   string
}

// GetMountPointData returns statfs information for each mount point in a channel.
func GetMountPointData() <-chan *MountPointData {
	out := make(chan *MountPointData)
	go func() {
		for _, mountPoint := range Config.Check.Mountpoint {
			statfs, err := MountPointStatus(mountPoint)
			if err != nil {
				Logger.Printf("Unable to get statfs data for mount point %q: %v", mountPoint, err)
			}
			out <- &MountPointData{percentAvail: PercentAvailable(statfs), mountPoint: mountPoint}
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
func PercentAvailable(statfs *syscall.Statfs_t) float32 {
	if statfs.Blocks == 0 {
		return 0
	}
	return float32(float64(statfs.Bavail) / float64(statfs.Blocks) * 100.0)
}
