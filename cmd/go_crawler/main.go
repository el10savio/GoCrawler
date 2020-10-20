package main

// GoCrawler main.go insantaties the GoCrawler worker node net/http server
// and establishes a connection to the RabbitMQ queue. It acts as a work queue
// node sending & receiving async crawl requests to the RabbitMQ message bus
// and persisting its crawled links in the Postgres Database. http requests
// a URL can also be sent directly to the node to process along with messages

import (
	"github.com/el10savio/GoCrawler/handlers"
	"github.com/el10savio/GoCrawler/internal/parser"
	"github.com/el10savio/GoCrawler/internal/platform/database"
	messagebus "github.com/el10savio/GoCrawler/internal/platform/message_bus"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	// PORT is the GoCrawler
	// http server port
	// Todo:  we can put this in env variable
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {

	// generates the routes todo: this can be encapsulated again
	var routes = []handlers.Route{
		{"/", "GET", handlers.CrawlerIndex},
		{"/crawler/parse", "POST", handlers.ParseHandler},
	}

	// Initialize the server routes
	r := handlers.Router(routes)

	// Initialize the postgres database
	db, err := database.Initialize()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Info("successfully connected to postgres database")

	// Connect to RabbitMQ
	connection, err := messagebus.Connect()
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	log.Info("successfully connected to rabbitmq messagebus")

	channel, err := messagebus.CreateChannel(connection)
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	log.Info("successfully created rabbitmq channel")

	err = messagebus.EstablishPublishQueue(channel)
	if err != nil {
		panic(err)
	}

	log.Info("successfully created rabbitmq publish queue")

	messages, err := messagebus.ConsumeMessages("links", channel)
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
