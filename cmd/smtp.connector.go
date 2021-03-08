package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"

	"github.com/JBGoldberg/uhura/models"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
)

func sendSMTP(_email models.Email) error {
	log.Printf("Sending message using SMTP server @ %s:%d", config.smtp.serverHost, config.smtp.serverPort)

	if len(_email.ID) == 0 {

		id, err := uuid.NewV4()
		if err != nil {
			return err
		}

		_email.ID = id.String()

	}
	log.Printf("Sending message %s from %s, about %s", _email.ID, _email.From, _email.Subject)

	// conn, err := net.Dial("tcp", "workcluster.nekutima.eu:25")
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.smtp.serverHost, config.smtp.serverPort))
	if err != nil {
		return err
	}

	// c, err := smtp.NewClient(conn, "ip5f5bf741.dynamic.kabel-deutschland.de")
	c, err := smtp.NewClient(conn, config.smtp.clientHost)
	if err != nil {
		return err
	}
	defer c.Close()

	if err = c.Hello(config.smtp.clientHost); err != nil {
		return err
	}

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: config.smtp.serverHost}
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}

	if err = c.Mail(_email.From); err != nil {
		return err
	}

	if err = c.Rcpt(_email.To); err != nil {
		return err
	}

	for _, bcc := range _email.Bcc {
		if err = c.Rcpt(bcc); err != nil {
			return err
		}
	}

	for _, cc := range _email.Cc {
		if err = c.Rcpt(cc); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	var body string
	body += fmt.Sprintf("To: %s\r\n", _email.To)
	body += fmt.Sprintf("From: %s\r\n", _email.From)
	body += fmt.Sprintf("Subject: %s\r\n", _email.Subject)
	body += fmt.Sprintf("Uhura-ID: %s\r\n", _email.ID)

	for _, bcc := range _email.Bcc {
		body += fmt.Sprintf("Bcc: %s\r\n", bcc)
	}

	for _, cc := range _email.Cc {
		body += fmt.Sprintf("Cc: %s\r\n", cc)
	}

	body += fmt.Sprintf("\r\n%s\r\n", _email.Message)

	if _, err = w.Write([]byte(body)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	if err = c.Quit(); err != nil {
		return err
	}

	return nil
}
