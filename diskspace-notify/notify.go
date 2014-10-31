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

// Creates notification event. These events will be collected and sent in one email.
func GetEvent(percentAvail float32, mountPoint string) *Event {
	return &Event{percentAvail: percentAvail, mountPoint: mountPoint}
}

// Collects all notification events and sends them according to config settings.
func SendNotification(eventQueue chan *Event) error {
	// Do nothing without data in Queue.
	if len(eventQueue) == 0 {
		return nil
	}

	var err error

	auth := smtp.PlainAuth(
		"",
		Config.Smtp.Username,
		Config.Smtp.Password,
		Config.Smtp.Address,
	)
	serverAddr := fmt.Sprintf("%v:%d", Config.Smtp.Address, Config.Smtp.Port)

	body := ""

	for event := range eventQueue {
		body += fmt.Sprintf(
			Config.Mail.Message+"\n",
			event.mountPoint,
			event.percentAvail,
		)
	}

	message := fmt.Sprintf(
		emailTemplate,
		Config.Mail.From,
		strings.Join(Config.Mail.Sendto, ", "),
		Config.Mail.Subject,
		body,
	)

	err = smtp.SendMail(
		serverAddr,
		auth,
		Config.Mail.From,
		Config.Mail.Sendto,
		[]byte(message),
	)
	if err != nil {
		Logger.Printf("Error while sending email: %v", err)
	}
	Logger.Printf("Sent notification.")

	return err
}
