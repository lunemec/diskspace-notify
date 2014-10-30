package main

import (
	"syscall"
)

type MountPointData struct {
	percentAvail *float32
	mountPoint   string
}

func GetMountPointData() <-chan *MountPointData {
	var (
		percentAvail *float32
		err          error
	)
	out := make(chan *MountPointData)

	go func() {
		for _, mountPoint := range Config.Check.Mountpoint {
			fsStat, err = MountPointStatus(mountPoint)

			if err != nil {
				Logger.Printf("Unable to get Statfs data: %v", err)
			}

			percentAvail = PercentAvailable(fsStat)
			out <- &MountPointData{percentAvail: percentAvail, mountPoint: mountPoint}
		}
		close(out)
	}()
	return out
}

func MountPointStatus(mountpoint string) (*syscall.Statfs_t, error) {
	/*
		Checks specified mount-point and returns its status.
	*/
	var fsStats syscall.Statfs_t

	err := syscall.Statfs(mountpoint, &fsStats)
	return &fsStats, err
}

func PercentAvailable(fsStat *syscall.Statfs_t) *float32 {
	/*
		Calculates how many percent are available given input fsStat data.
	*/
	var (
		totalSize    uint64
		totalAvail   uint64
		percentAvail float32
	)

	totalSize = uint64(fsStat.Bsize) * fsStat.Blocks
	totalAvail = uint64(fsStat.Bsize) * fsStat.Bavail
	percentAvail = float32(float64(totalAvail) / (float64(totalSize) / float64(100)))

	return &percentAvail
}
