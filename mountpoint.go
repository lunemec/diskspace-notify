package main

import (
	"github.com/dustin/go-humanize"
	"sync"
	"syscall"
)

type MountPoint struct {
	percentAvail uint8
	freeSpace    string // Output from humanize.
	totalSize    string // Output frin humanize.
	mountPoint   string
}

var wg sync.WaitGroup

// mountPointStatus checks specified mount point and returns its status.
func mountPointStatus(mountpoint string) (*syscall.Statfs_t, error) {
	var statfs syscall.Statfs_t
	err := syscall.Statfs(mountpoint, &statfs)
	return &statfs, err
}

// percentAvailable calculates how many percent are available given input statfs data.
func percentAvailable(statfs *syscall.Statfs_t) uint8 {
	if statfs.Blocks == 0 {
		return uint8(0)
	}
	return uint8(float32(statfs.Bavail) / float32(statfs.Blocks) * 100.0)
}

// Checks if given value is lower or equal to Threshold.
func checkThreshold(value uint8) bool {
	if value <= Config.Check.Threshold {
		return true
	}
	return false
}

// Checks data for single mountpoint.
func CheckMountPoint(c chan *MountPoint, mountPoint string) {
	defer wg.Done()
	statfs, err := mountPointStatus(mountPoint)

	if err != nil {
		Logger.Printf("Unable to get statfs data for mount point %q: %v.", mountPoint, err)
	}

	// We want to check threshold before creating data so we don't waste resources.
	percent := percentAvailable(statfs)
	if checkThreshold(percent) {
		c <- &MountPoint{
			percentAvail: percent,
			freeSpace:    humanize.Bytes(statfs.Bavail * uint64(statfs.Bsize)),
			totalSize:    humanize.Bytes(statfs.Blocks * uint64(statfs.Bsize)),
			mountPoint:   mountPoint,
		}
	}
}

// MountPointData returns statfs information for each mount point in a channel.
func MountPointData() []*MountPoint {
	mountPoints := []*MountPoint{}
	c := make(chan *MountPoint, len(Config.Check.Mountpoint))

	// Spawn goroutine for each mountpoint.
	for _, mountPoint := range Config.Check.Mountpoint {
		wg.Add(1)
		go CheckMountPoint(c, mountPoint)
	}
	// Wait for all goroutines to complete and close channel.
	wg.Wait()
	close(c)

	for data := range c {
		mountPoints = append(mountPoints, data)
	}
	return mountPoints
}
