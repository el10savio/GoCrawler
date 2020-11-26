package main

// GoCrawler main.go insantaties the GoCrawler worker node net/http server
// and establishes a connection to the RabbitMQ queue. It acts as a work queue
// node sending & receiving async crawl requests to the RabbitMQ message bus
// and persisting its crawled links in the Postgres Database. http requests
// a URL can also be sent directly to the node to process along with messages

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/GoCrawler/GoCrawler/database"
	"github.com/el10savio/GoCrawler/GoCrawler/handlers"
	"github.com/el10savio/GoCrawler/GoCrawler/messageBus"
	"github.com/el10savio/GoCrawler/GoCrawler/parser"
)

const (
	// PORT is the GoCrawler
	// http server port
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Initialize the server routes
	r := handlers.Router()

	// Initialize the postgres database
	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Info("successfully connected to postgres database")

	// Connect to RabbitMQ
	connection, err := messageBus.Connect()
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	log.Info("successfully connected to rabbitmq messagebus")

	channel, err := messageBus.CreateChannel(connection)
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	log.Info("successfully created rabbitmq channel")

	err = messageBus.EstablishPublishQueue(channel)
	if err != nil {
		panic(err)
	}

	log.Info("successfully created rabbitmq publish queue")

	messages, err := messageBus.ConsumeMessages("links", channel)
	if err != nil {
		panic(err)
	}

	log.Info("successfully subscribed to rabbitmq links queue")

	go func() {
		for delivery := range messages {
			URL := string(delivery.Body)
			log.WithFields(log.Fields{"url": URL}).Info("received message")
			parser.Parse(URL)
		}
	}()

	// Start the HTTP server
	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started gocrawler node server")

	http.ListenAndServe(":"+PORT, r)
}
