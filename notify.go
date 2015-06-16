package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

const emailTemplate = `From: %v
To: %v
Subject: %v

%v
`

// Collects all notification events and sends them according to config settings.
func SendNotification(mountPoints []*MountPoint) error {
	var err error

	var auth smtp.Auth
	if Config.Smtp.Auth {
		auth = smtp.PlainAuth(
			"",
			Config.Smtp.Username,
			Config.Smtp.Password,
			Config.Smtp.Address,
		)
	} else {
		auth = nil
	}
	serverAddr := fmt.Sprintf("%v:%d", Config.Smtp.Address, Config.Smtp.Port)

	body := ""

	for _, data := range mountPoints {
		body += fmt.Sprintf(
			Config.Mail.Message+"\n",
			data.mountPoint,
			data.percentAvail,
			data.freeSpace,
			data.totalSize,
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
		Logger.Printf("Error while sending email: %v.", err)
	} else {
		Logger.Printf("Sent notification.")
	}
	return err
}
