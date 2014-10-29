package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Event struct {
	percentAvail float32
	mountPoint   string
}

const emailTemplate = `From: %v
To: %v
Subject: %v

%v
`

func GetEvent(percentAvail float32, mountPoint string) *Event {
	/*
		Creates notification event. These events will be collected and sent in one email.
	*/
	return &Event{percentAvail: percentAvail, mountPoint: mountPoint}
}

func SendNotification(config *Config, eventQueue chan *Event) error {
	/*
		Collects all notification events and sends them according to config settings.
	*/
	var err error

	auth := smtp.PlainAuth(
		"",
		config.Smtp.Username,
		config.Smtp.Password,
		config.Smtp.Address,
	)
	serverAddr := fmt.Sprintf("%v:%d", config.Smtp.Address, config.Smtp.Port)

	body := ""

	for event := range eventQueue {
		body += fmt.Sprintf(
			config.Mail.Message+"\n",
			event.mountPoint,
			event.percentAvail,
		)
	}

	message := fmt.Sprintf(
		emailTemplate,
		config.Mail.From,
		strings.Join(config.Mail.Sendto, ", "),
		config.Mail.Subject,
		body,
	)

	err = smtp.SendMail(
		serverAddr,
		auth,
		config.Mail.From,
		config.Mail.Sendto,
		[]byte(message),
	)
	if err != nil {
		fmt.Printf("Error while sending email: %v", err)
	}

	return err
}
