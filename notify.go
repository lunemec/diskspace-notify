package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Event struct {
	eventData *MountPoint
}

const emailTemplate = `From: %v
To: %v
Subject: %v

%v
`

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
			event.eventData.mountPoint,
			event.eventData.percentAvail,
			event.eventData.freeSpace,
			event.eventData.totalSize,
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
