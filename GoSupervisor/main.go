package main

// GoSupervisor main.go insantaties the GoSupervisor net/http server
// and establishes a connection to the RabbitMQ queue. It acts as a web
// crawler proxy sending async crawl requests for each GoCrawler worker
// node to process and aggregates the resultant crawled links for the client

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/GoCrawler/GoSupervisor/database"
	"github.com/el10savio/GoCrawler/GoSupervisor/handlers"
	"github.com/el10savio/GoCrawler/GoSupervisor/messageBus"
)

const (
	// PORT is the GoSupervisor
	// http server port
	PORT = "8050"
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

	// Start the HTTP server
	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started gosupervisor node server")

	http.ListenAndServe(":"+PORT, r)
}
