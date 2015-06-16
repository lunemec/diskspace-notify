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
		Auth          bool
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

const configError string = "Config file rror: "

// Default config values.
const defaultPort int = 25
const defaultAntiSpamDelay int = 3600
const defaultDelay int = 10
const defaultThreshold uint8 = uint8(10)

// Returns `true` if all items in array are true.
func all(values []bool) bool {
	for _, value := range values {
		if !value {
			return false
		}
	}

	return true
}

func checkIntField(configValue int, nullValue int, errorMsg string, logit bool) bool {
	if configValue == nullValue {
		if logit {
			Logger.Printf("%v%v\n", configError, errorMsg)
		}
		return false
	}

	return true
}

func checkUint8Field(configValue uint8, nullValue uint8, errorMsg string, logit bool) bool {
	if configValue == nullValue {
		if logit {
			Logger.Printf("%v%v\n", configError, errorMsg)
		}
		return false
	}

	return true
}

func checkStringField(configValue string, nullValue string, errorMsg string, logit bool) bool {
	if configValue == nullValue {
		if logit {
			Logger.Printf("%v%v\n", configError, errorMsg)
		}
		return false
	}

	return true
}

func checkStringArrayField(configValue []string, length int, errorMsg string, logit bool) bool {
	if len(configValue) == length {
		if logit {
			Logger.Printf("%v%v\n", configError, errorMsg)
		}
		return false
	}

	return true
}

// Checks required items in config and returns errors.
func checkRequired(config *ConfigData) error {
	ok := []bool{}

	// [mail] section check.
	ok = append(ok, checkStringField(config.Mail.From, "", "[mail] `from` is required.", true))
	ok = append(ok, checkStringArrayField(config.Mail.Sendto, 0, "[mail] `sendto` is required.", true))
	ok = append(ok, checkStringField(config.Mail.Subject, "", "[mail] `subject` is required.", true))
	ok = append(ok, checkStringField(config.Mail.Message, "", "[mail] `message` is required.", true))

	// [smtp] section check.
	ok = append(ok, checkStringField(config.Smtp.Address, "", "[smtp] `address` is required.", true))
	// Check Username only when SMTP Auth is enabled. Password may be empty.
	if config.Smtp.Auth {
		ok = append(ok, checkStringField(config.Smtp.Username, "", "[smtp] `username` is required.", true))
	}

	// [check] section check.
	ok = append(ok, checkStringArrayField(config.Check.Mountpoint, 0, "[check] `mountpoint` is required.", true))

	if !all(ok) {
		return fmt.Errorf("some required config items is missing.")
	}
	return nil
}

// Check if not-required values are set, and if not, fill in defaults.
func checkDefaults(config *ConfigData) {
	ok := checkIntField(config.Smtp.Port, 0, "", false)
	if !ok {
		config.Smtp.Port = defaultPort
	}

	ok = checkIntField(config.Smtp.AntiSpamDelay, 0, "", false)
	if !ok {
		config.Smtp.AntiSpamDelay = defaultAntiSpamDelay
	}

	ok = checkIntField(config.Check.Delay, 0, "", false)
	if !ok {
		config.Check.Delay = defaultDelay
	}

	ok = checkUint8Field(config.Check.Threshold, uint8(0), "", false)
	if !ok {
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
