package spider

import (
	"../messageBus"
	"github.com/streadway/amqp"
)

var (
	// Instantiate a shared RabbitMQ channel
	// variable to be shared in the package
	channel *amqp.Channel
)

func init() {
	// Connect to RabbitMQ
	connection, err := messageBus.Connect()
	if err != nil {
		panic(err)
	}

	// Create and establish RabbitMQ
	// connection channel
	channel, err = messageBus.CreateChannel(connection)
	if err != nil {
		panic(err)
	}
}

// Publish link message to RabbitMQ
func PublishLink(link string) error {
	// Create message from the given link
	message := messageBus.CreateMessage(link)

	// Publish message to RabbitMQ
	err := messageBus.PublishMessage(message, channel)
	if err != nil {
		return err
	}

	return nil
}
