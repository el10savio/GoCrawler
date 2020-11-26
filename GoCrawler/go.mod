module github.com/el10savio/GoCrawler/GoCrawler

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.8.0
	github.com/sirupsen/logrus v1.7.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
)

replace github.com/el10savio/GoCrawler/GoCrawler/database => ../database

replace github.com/el10savio/GoCrawler/GoCrawler/messageBus => ../messageBus

replace github.com/el10savio/GoCrawler/GoCrawler/parser => ../parser
