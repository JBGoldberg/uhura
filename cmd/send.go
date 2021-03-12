package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JBGoldberg/uhura/smtp"

	"github.com/JBGoldberg/uhura/messaging"
	"github.com/JBGoldberg/uhura/models"
	"github.com/ThreeDotsLabs/watermill/message"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send smtp|telegram",
	Short: "Start process(es) to send messages",
	Long:  `Start the processes that wait for messages in queues and send them using the respective protocols.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Sending messages")

		for _, c := range args {

			switch c {
			case "smtp":
				if err := checkSMTPQueue(); err != nil {
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

func checkSMTPQueue() error {

	server := fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		config.ampq.username,
		config.ampq.password,
		config.ampq.serverHost,
		config.ampq.serverPort)

	subscriber, err := messaging.NewSubscriber(server)
	if err != nil {
		return err
	}

	SMTPQeue, err := subscriber.Subscribe(context.Background(), config.ampq.queues.smtp)
	if err != nil {
		return err
	}

	return processSMTPQueue(SMTPQeue)
}

func processSMTPQueue(messages <-chan *message.Message) error {

	emailServer, err := smtp.NewServer(config.smtp.serverHost, config.smtp.serverPort, config.smtp.clientHost)
	if err != nil {
		log.Error("Unable to connect to SMTP server")
		return err
	}

	for msg := range messages {

		var email models.Email
		err := json.Unmarshal([]byte(msg.Payload), &email)
		if err != nil {
			log.Error("Unable to deJSON message", err, msg.UUID)
			msg.Reject()
			continue
		}

		if err := emailServer.SendSMTP(email); err != nil {
			log.Error("Unable to send via smtp email", err, email.ID)
			msg.Nack()
			continue
		}

		log.Printf("SMTP email sent as requested by: %s", msg.UUID)
		msg.Ack()
	}

	return nil
}

// func processSMTPQueue(emails []models.Email) error {
// 	log.Printf("Processing SMTP queue using server @ %s:%d", config.smtp.serverHost, config.smtp.serverPort)

// 	var counter int
// 	for _, email := range emails {
// 		if err := sendSMTP(email); err != nil {
// 			return err
// 		}
// 		counter++
// 	}

// 	log.Printf("SMTP queue finished with %d emails sent", counter)
// 	return nil
// }

//func (svc *sservice) process(messages <-chan *message.Message) {
//
//	for msg := range messages {
//
//		var email Email
//		json.Unmarshal(msg.Payload, &email)
//		if err := svc.sendTransationalEmail(email); err != nil {
//			log.Println(err)
//			msg.Nack()
//			continue
//		}
//		msg.Ack()
//	}
//
//}
