package main

import (
	"code.google.com/p/gcfg"
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


// Checks required items in config and returns errors.
func checkRequired(config *ConfigData) error {

}

// Loads configuration file located under `configFile` path
// and returns pointer to it.
func LoadConfig(configFile string) (*ConfigData, error) {
	var config ConfigData
	min := uint8(0)
	max := uint8(100)

	err := gcfg.ReadFileInto(&config, configFile)
	if err != nil {
	
	}
	err := checkRequired(&config)
	err := 

	if config.Check.Threshold <= min || config.Check.Threshold >= max {
		err = fmt.Errorf("Wrong config value, threshold must be larger than 0 and lower than 100.")
	}

	return &config, err
}
