package cmd

import (
	"errors"

	"github.com/JBGoldberg/uhura/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var email = models.Email{
	ID:   "teste-1234",
	From: "uhura@nekutima.eu",
	To:   "jimbrunogoldberg@gmail.com",
	Cc: []string{
		"jim@bycoders.co",
	},
	Bcc: []string{
		"jbgoldberg@nekutima.eu",
	},
	Subject: "Test Email From Uhura",
	Message: "This is a message file from Uhura!",
}

var sendCmd = &cobra.Command{
	Use:   "send smtp|telegram",
	Short: "Send messages on queue",
	Long:  `Reads the messages queue to be send and process it.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Sending messages")

		for _, c := range args {

			switch c {
			case "smtp":

				emailq := []models.Email{
					email,
				}

				if err := processSMTPQueue(emailq); err != nil {
					return err
				}
				break

			case "telegram":
				log.Errorf("Telegram communications are not implemmented yet")
				return errors.New("not implemented")

			default:
				log.Errorf("I don't know the %s communication channel", c)
				return errors.New("channel unknown")

			}
		}

		return nil

	},
}
