package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"
)

var (
	configFile    = flag.String("config", "", "Path to config file.")
	config        *Config
	emailSent     = false
	emailSendTime time.Time
	fsStat        *syscall.Statfs_t
)

func main() {
	// TODO errors should be able to be logged to a file (flag parameter) or to stdout (default)
	// TODO daemonize (flag parameter)
	// Parse command line arguments.
	flag.Parse()
	if *configFile == "" {
		fmt.Printf("Config argument required. \n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var (
		event        *Event
		mountPoint   string
		percentAvail *float32
		err          error
	)

	fmt.Println("Free disk space check.")
	config, err = LoadConfig(*configFile)

	if err != nil {
		fmt.Printf("Can't load configuration file %s, error: %s\n", *configFile, err)
		os.Exit(1)
	}

	for {
		// buffered event queue
		EventQueue := make(chan *Event, 100)

		for mountPointData := range GetMountPointData(config) {
			percentAvail = mountPointData.percentAvail
			mountPoint = mountPointData.mountPoint

			if *percentAvail > config.Check.Threshold {
				event = GetEvent(*percentAvail, mountPoint)
				EventQueue <- event
			}
		}
		close(EventQueue)

		// Anti-spam measure.
		if emailSent {
			if time.Since(emailSendTime) >= time.Duration(config.Smtp.AntiSpamDelay)*time.Second {
				err = SendNotification(config, EventQueue)
				emailSent = true
			}
		} else {
			err = SendNotification(config, EventQueue)
			emailSent = true
			emailSendTime = time.Now()
		}

		if err != nil {
			fmt.Printf("Unable to send email: %v", err)
		}

		time.Sleep(time.Duration(config.Check.Delay) * time.Second)
	}
}
