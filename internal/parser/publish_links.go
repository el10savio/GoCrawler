package parser

import (
	messageBus "github.com/el10savio/GoCrawler/internal/platform/message_bus"
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

// Publish links messages to RabbitMQ
func PublishLinks(links []string) error {
	if len(links) == 0 {
		return nil
	}

	for _, link := range links {
		// Create message from the given link
		message := messageBus.CreateMessage(link)

		// Publish message to RabbitMQ
		err := messageBus.PublishMessage(message, channel)
		if err != nil {
			return err
		}
	}

	return nil
}
