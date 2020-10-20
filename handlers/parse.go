package handlers

import (
	"encoding/json"
	"github.com/el10savio/GoCrawler/internal/parser"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ParseHandler is the http handler for /crawler/parse to process
// the crawl request and parse the URL and obtain all the links in it
func ParseHandler(w http.ResponseWriter, r *http.Request) {
	// Get URL from JSON body
	URL, err := GetURL(r)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("unable to parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Parse Obtained URL
	links, err := parser.Parse(URL)
	if err != nil {
		log.WithFields(log.Fields{
			"URL":   URL,
			"error": err,
		}).Error("unable to parse given URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log indicating
	// successful URL parse
	log.WithFields(log.Fields{
		"URL": URL,
	}).Debug("successful URL parse")

	// json encode links response
	json.NewEncoder(w).Encode(links)
}
