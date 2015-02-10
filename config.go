package main

import (
	"code.google.com/p/gcfg"
	"errors"
	"fmt"
)

type ConfigData struct {
	Mail struct {
		From    string
		Sendto  []string
		Subject string
		Message string
	}
	Smtp struct {
		Address       string
		Port          int
		Username      string
		Password      string
		AntiSpamDelay int
	}
	Check struct {
		Delay      int
		Threshold  uint8
		Mountpoint []string
	}
}

type errorlist []error

const configError string = "Config file rror: "

// Default config values.
const defaultPort int = 25
const defaultAntiSpamDelay int = 3600
const defaultDelay int = 10
const defaultThreshold uint8 = uint8(10)

func (e *errorlist) Add(err string) {
	*e = append(*e, errors.New(err))
}

// Checks required items in config and returns errors.
func checkRequired(config *ConfigData) error {
	var errs errorlist

	// [mail] section check.
	if config.Mail.From == "" {
		errs.Add("[mail] `from` is required.")
	}
	if len(config.Mail.Sendto) == 0 {
		errs.Add("[mail] `sendto` is required.")
	}
	if config.Mail.Subject == "" {
		errs.Add("[mail] `subject` is required.")
	}
	if config.Mail.Message == "" {
		errs.Add("[mail] `message` is required.")
	}

	// [smtp] section check.
	if config.Smtp.Address == "" {
		errs.Add("[smtp] `address` is required.")
	}
	if config.Smtp.Username == "" {
		errs.Add("[smtp] `username` is required.")
	}
	if config.Smtp.Password == "" {
		errs.Add("[smtp] `password` is required.")
	}

	// [check] section check.
	if len(config.Check.Mountpoint) == 0 {
		errs.Add("[check] `mountpoint` is required.")
	}

	if len(errs) != 0 {
		for _, err := range errs {
			Logger.Printf("%v\n", err)
		}
		return fmt.Errorf("some required config items is missing.")
	}
	return nil
}

// Check if not-required values are set, and if not, fill in defaults.
func checkDefaults(config *ConfigData) {
	if config.Smtp.Port == 0 {
		config.Smtp.Port = defaultPort
	}
	if config.Smtp.AntiSpamDelay == 0 {
		config.Smtp.AntiSpamDelay = defaultAntiSpamDelay
	}
	if config.Check.Delay == 0 {
		config.Check.Delay = defaultDelay
	}
	if config.Check.Threshold == uint8(0) {
		config.Check.Threshold = defaultThreshold
	}
}

// Loads configuration file located under `configFile` path
// and returns pointer to it.
func LoadConfig(configFile string) (*ConfigData, error) {
	var config ConfigData
	min := uint8(0)
	max := uint8(100)

	err := gcfg.ReadFileInto(&config, configFile)
	if err != nil {
		Logger.Fatalf("%v%v.\n", configError, err)
	}

	err = checkRequired(&config)
	if err != nil {
		Logger.Fatalf("%v%v\n", configError, err)
	}

	checkDefaults(&config)

	if config.Check.Threshold <= min || config.Check.Threshold >= max {
		err = fmt.Errorf("%v[check] threshold must be larger than 0 and lower than 100.\n", configError)
	}

	return &config, err
}
