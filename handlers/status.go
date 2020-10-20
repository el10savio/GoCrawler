package handlers

import (
	"encoding/json"
	"github.com/el10savio/GoCrawler/internal/spider"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Status is the http handler for /spider/view to process
// the given spider view request and aggregate the crawled
// URLs for the client to view as the JSON response
func Status(w http.ResponseWriter, r *http.Request) {
	// Get URL from JSON body
	URL, err := GetURL(r)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("unable to parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Status Of Obtained URL
	sitemap, err := spider.StatusLink(URL)
	if err != nil {
		log.WithFields(log.Fields{
			"URL":   URL,
			"error": err,
		}).Error("unable to obtain status of given URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log indicating
	// successful URL status
	log.WithFields(log.Fields{
		"URL": URL,
	}).Debug("successful URL status")

	// json encode result sitemap
	json.NewEncoder(w).Encode(sitemap)
}
