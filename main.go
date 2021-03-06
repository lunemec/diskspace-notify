package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

var (
	configFile    = flag.String("config", "", "Path to config file.")
	logFilePath   = flag.String("log", "", "Path to log file.")
	defaultConfig = flag.Bool("defaultconfig", false, "Prints default config to stdout.")
	printVersion  = flag.Bool("version", false, "Print current version.")

	emailSendTime time.Time
	err           error
	fsStat        *syscall.Statfs_t

	Config *ConfigData
	Logger *log.Logger
)

const version = "Diskspace-notify 1.5 \n"

// Initializes log depending on logFilePath variable.
// If it is empty, it will log to `stderr`.
// If not, it will create/append to a file on that path.
func initLog(logFilePath *string) {
	logFile := os.Stderr

	if *logFilePath != "" {
		f, err := os.OpenFile(*logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Can't open file for logging, error: %v.\n", err)
		}
		logFile = f
	}

	Logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

// Initializes configuration file under path `configFile`.
func initConfig(configFile *string) {
	Config, err = LoadConfig(*configFile)
	if err != nil {
		Logger.Fatalf("Can't load configuration file %v, error: %s.\n", *configFile, err)
	}
}

// Sends Email only after antiSpamDelay has passed (3600s default)
// after the email was sent for the first time.
func sendEmail(mountPoints []*MountPoint, emailSent bool, antiSpamDelay time.Duration) bool {
	if emailSent {
		waiting := uint64((antiSpamDelay - time.Since(emailSendTime)).Seconds())
		Logger.Printf("Email already sent. Waiting %vs before next email.\n", waiting)

		if time.Since(emailSendTime) >= antiSpamDelay {
			err = SendNotification(mountPoints)
			if err != nil {
				Logger.Printf("Error while sending email: %v.\n", err)
			} else {
				emailSent = true
				emailSendTime = time.Now()
			}
		}
	} else {
		err = SendNotification(mountPoints)
		if err != nil {
			Logger.Printf("Error while sending email: %v.\n", err)
		} else {
			emailSent = true
			emailSendTime = time.Now()
		}
	}

	return emailSent
}

func main() {

	var (
		mountPoints []*MountPoint
		emailSent   = false
	)

	// Parse command line arguments.
	flag.Parse()
	if *printVersion {
		fmt.Printf("%v", version)
		os.Exit(0)
	}
	if *defaultConfig {
		PrintDefaultConfig()
		os.Exit(0)
	}
	if *configFile == "" {
		log.Printf("Config argument required. \n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	initLog(logFilePath)
	initConfig(configFile)

	Logger.Printf("Starting ...")

	// Set delays, these must be set after config is initialized.
	antiSpamDelay := time.Duration(Config.Smtp.AntiSpamDelay) * time.Second
	sleepDelay := time.Duration(Config.Check.Delay) * time.Second

	for {
		mountPoints = MountPointData()
		if len(mountPoints) > 0 {
			emailSent = sendEmail(mountPoints, emailSent, antiSpamDelay)
		}

		time.Sleep(sleepDelay)
	}
}
