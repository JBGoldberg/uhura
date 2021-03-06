package cmd

import (
	"github.com/JBGoldberg/uhura/models"
	log "github.com/sirupsen/logrus"
)

func processSMTPQueue(emails []models.Email) error {
	log.Printf("Processing SMTP queue using server @ %s:%d", config.smtp.serverHost, config.smtp.serverPort)

	var counter int
	for _, email := range emails {
		if err := sendSMTP(email); err != nil {
			return err
		}
		counter++
	}

	log.Printf("SMTP queue finished with %d emails sent", counter)
	return nil
}
