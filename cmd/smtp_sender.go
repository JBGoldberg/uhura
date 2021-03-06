package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

func smtpSend() error {
	log.Printf("Sending message using SMTP server @ %s:%d", config.smtp.serverHost, config.smtp.serverPort)

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

	if err = c.Mail("oompa@nekutima.eu"); err != nil {
		return err
	}

	to := []string{
		"jimbrunogoldberg@gmail.com",
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("To: jimbrunogoldberg@gmail.com\r\n" +
		"From: uhura@nekutima.eu\r\n" +
		"Subject: Test Email From Uhura\r\n" +
		"\r\n" +
		"This is some content to reach a user...\r\n"))
	if err != nil {
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
