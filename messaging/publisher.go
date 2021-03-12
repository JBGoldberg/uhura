package messaging

import (
	log "github.com/sirupsen/logrus"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

// Connect to AMQP server to publish on queues
func NewPublisher(amqpConnectString string) (*amqp.Publisher, error) {

	log.Infof("Retreiving publisher from %s", removePassword(amqpConnectString))

	amqpConfig := amqp.NewDurableQueueConfig(amqpConnectString)
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

// publishMessages(publisher)
