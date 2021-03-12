package cmd

import (
	"encoding/json"
	"errors"

	"github.com/JBGoldberg/uhura/messaging"
	"github.com/JBGoldberg/uhura/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

func init() {

	requestCmd.PersistentFlags().StringVarP(&emailr.Subject, "subject", "s", "", "subject of the message")
	requestCmd.PersistentFlags().StringVarP(&emailr.To, "to", "t", "", "recepient of the message")
	requestCmd.PersistentFlags().StringVarP(&emailr.From, "from", "f", "", "sender of the message")
	requestCmd.PersistentFlags().StringVarP(&emailr.Message, "message", "m", "", "conteudo da mensagem")

	rootCmd.AddCommand(requestCmd)
}

var emailr = models.Email{
	Cc: []string{
		"jim@bycoders.co",
	},
	Bcc: []string{
		"jbgoldberg@nekutima.eu",
	},
	Message: "One more pipe in place from Uhura!",
}

var requestCmd = &cobra.Command{
	Use:   "request smtp|telegram",
	Short: "Request messaging through a channel.",
	Long:  `Request messaging an address using a channel or channels`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Requesting message")

		// TODO Get message from stdin

		for _, c := range args {

			switch c {
			case "smtp":
				if err := requestSMTPQueue(emailr); err != nil {
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

func requestSMTPQueue(_email models.Email) error {

	server := fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		config.ampq.username,
		config.ampq.password,
		config.ampq.serverHost,
		config.ampq.serverPort)

	publisher, err := messaging.NewPublisher(server)
	if err != nil {
		log.Error("Unable to get publisher")
		return err
	}

	json, err := json.Marshal(emailr)
	if err != nil {
		log.Error("Unable to JSONed Email")
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte(json))

	err = publisher.Publish(config.ampq.queues.smtp, msg)
	if err != nil {
		log.Error("Unable queue message")
		return err
	}

	return nil

}
