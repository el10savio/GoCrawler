package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/GoCrawler/GoSupervisor/spider"
)

// Publish is the http handler for /spider/crawl to process
// the given spider crawl request and publish the URL
// to RabbitMQ for the GoCrawler nodes to process
func Publish(w http.ResponseWriter, r *http.Request) {
	// Get URL from JSON body
	URL, err := GetURL(r)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("unable to parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Publish Obtained URL
	err = spider.PublishLink(URL)
	if err != nil {
		log.WithFields(log.Fields{
			"URL":   URL,
			"error": err,
		}).Error("unable to publish given URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log indicating
	// successful URL publish
	log.WithFields(log.Fields{
		"URL": URL,
	}).Debug("successful URL publish")

	// json encode success
	w.WriteHeader(http.StatusOK)
}
