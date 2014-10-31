package main

import (
	"code.google.com/p/gcfg"
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
		Threshold  float32
		Mountpoint []string
	}
}

// Loads configuration file located under `configFile` path
// and returns pointer to it.
func LoadConfig(configFile string) (*ConfigData, error) {
	var config ConfigData

	err := gcfg.ReadFileInto(&config, configFile)
	return &config, err
}
