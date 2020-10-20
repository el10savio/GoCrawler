package parser

import (
	"github.com/el10savio/GoCrawler/internal/platform/database"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

// Parse is the main function that takes in a given
// URL and obtains all the urls present in it
func Parse(URL string) ([]string, error) {
	// Exit if already parsed before
	count, err := database.GetParentCount(URL)
	if err != nil || count != 0 {
		return []string{}, err
	}

	// Validate the given URL
	err = ValidateURL(URL)
	if err != nil {
		_ = database.RemoveLink(URL)
		return []string{}, err
	}

	log.WithFields(log.Fields{"URL": URL}).Info("successfully validated url")

	// Get Links
	links, err := GetLinks(URL)
	if err != nil {
		return []string{}, err
	}

	log.WithFields(log.Fields{"URL": URL}).Info("successfully obtained links from url")

	// Persist Links
	err = InsertLinks(URL, links)
	if err != nil {
		return []string{}, err
	}

	log.WithFields(log.Fields{"URL": URL}).Info("successfully inserted links from url into DB")

	// Publish Links
	err = PublishLinks(links)
	if err != nil {
		return []string{}, err
	}

	log.WithFields(log.Fields{"URL": URL}).Info("successfully published links from url to rabbitMQ")

	return links, nil
}

// ValidateURL checks if the
// URL is semantically valid
func ValidateURL(URL string) error {
	if URL == "" {
		return ErrEmptyURL
	}

	url, err := url.ParseRequestURI(URL)
	if err != nil {
		return ErrInvalidURL
	}

	if url.Host == "" {
		return ErrEmptyURLHost
	}

	return nil
}

// GetLinks sends an http GET request to the given URL
// to obtain its resultant html body & then parse it
func GetLinks(URL string) ([]string, error) {
	var client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// Send GET Request
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return []string{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return []string{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []string{}, err
	}

	// Get links
	links := GetHyperlinks(URL, response)

	// Filter obtained links
	links = FilterHyperlinks(URL, links)

	return links, nil
}

// InsertLinks persists our obtained links to the Db
func InsertLinks(URL string, links []string) error {
	for _, link := range links {
		err := database.InsertLink(URL, link)
		if err != nil {
			return err
		}
	}
	return nil
}
