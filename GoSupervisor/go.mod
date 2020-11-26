module github.com/el10savio/GoCrawler/GoSupervisor

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.8.0
	github.com/sirupsen/logrus v1.7.0
	github.com/streadway/amqp v1.0.0
)

replace github.com/el10savio/GoCrawler/GoSupervisor/database => ../database

replace github.com/el10savio/GoCrawler/GoSupervisor/messageBus => ../messageBus

replace github.com/el10savio/GoCrawler/GoSupervisor/parser => ../parser
