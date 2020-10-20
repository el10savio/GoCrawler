package message_bus

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

const (
	host     = "localhost"
	port     = 5672
	user     = "guest"
	password = "guest"
)

// Connect establishes a connection
// to the RabbitMQ instance
func Connect() (*amqp.Connection, error) {
	var url string

	if os.Getenv("RABBITMQ_HOST") != "" {
		url = fmt.Sprintf("amqp://%s:%s@%s", user, password, os.Getenv("RABBITMQ_HOST"))
	} else {
		url = fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	}

	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// CreateChannel create a channel
// from the given connection
func CreateChannel(connection *amqp.Connection) (*amqp.Channel, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

// CreateExchange creates an events exchange
func CreateExchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare("events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}

// EstablishPublishQueue create exchange & links queue
func EstablishPublishQueue(channel *amqp.Channel) error {
	err := CreateExchange(channel)
	if err != nil {
		return err
	}

	err = CreateQueueAndBind(channel, "links")
	if err != nil {
		return err
	}

	return nil
}

// CreateQueueAndBind declares a queue
// and binds it to our channel
func CreateQueueAndBind(channel *amqp.Channel, queue string) error {
	// Create a queue
	_, err := channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// Bind the queue to the exchange
	err = channel.QueueBind(queue, "#", "events", false, nil)
	if err != nil {
		return err
	}

	return nil
}

// CreateMessage takes in our message body and
// creates a message to used sent in our queue
func CreateMessage(body string) amqp.Publishing {
	return amqp.Publishing{
		Body: []byte(body),
	}
}

// PublishMessage sends our RabbitMQ messages to the queue
func PublishMessage(message amqp.Publishing, channel *amqp.Channel) error {
	return channel.Publish("events",
		"key",
		false,
		false,
		message,
	)
}

// ConsumeMessages subscribes and waits for messages in the queue
func ConsumeMessages(queue string, channel *amqp.Channel) (<-chan amqp.Delivery, error) {
	messages, err := channel.Consume(queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return messages, err
}
