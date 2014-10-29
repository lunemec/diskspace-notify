package main

import (
	"fmt"
	"syscall"
)

const (
	configFile = "/home/lukas/torrentsync/Go/conf/notify.conf"
)

var (
	config *Config
	fsStat *syscall.Statfs_t
)

func main() {
	var (
		event        *Event
		mountPoint   string
		percentAvail *float32
		err          error
	)
	// buffered event queue
	EventQueue := make(chan *Event, 100)

	fmt.Println("Free disk space check.")
	config, err = LoadConfig(configFile)

	if err != nil {
		panic(fmt.Sprintf("Can't load configuration file %s, error: %s", configFile, err))
	}

	for mountPointData := range GetMountPointData(config) {
		percentAvail = mountPointData.percentAvail
		mountPoint = mountPointData.mountPoint

		if *percentAvail > config.Check.Threshold {
			event = GetEvent(*percentAvail, mountPoint)
			EventQueue <- event
		}
	}
	close(EventQueue)

	err = SendNotification(config, EventQueue)

	if err != nil {
		fmt.Printf("Unable to send email: %v", err)
	}
}
