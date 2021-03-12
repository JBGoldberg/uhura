package messaging

import (
	log "github.com/sirupsen/logrus"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

// NewSubscriber connects to AMQP server to check the queues
func NewSubscriber(amqpConnectString string) (*amqp.Subscriber, error) {

	// log.Infof("Checking for emails to send on %s from %s", server, config.ampq.queues.smtp)

	log.Infof("Retreiving subscriber from %s", removePassword(amqpConnectString))

	amqpConfig := amqp.NewDurableQueueConfig(amqpConnectString)
	subscriber, err := amqp.NewSubscriber(
		// This config is based on this example: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
		// It works as a simple queue.
		//
		// If you want to implement a Pub/Sub style service instead, check
		// https://watermill.io/docs/pub-sub-implementations/#amqp-consumer-groups
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}
