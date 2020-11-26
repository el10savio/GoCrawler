package spider

import (
	"github.com/el10savio/GoCrawler/GoSupervisor/database"
)

// Status struct is a tree like data structure
// that saves each crawled web link followed by
// all the web links in the web link page
type Status struct {
	URL   string   `json:"url"`
	Links []Status `json:"links"`
}

// StatusLink takes a given URL and aggregates all
// the crawled webpages recursively under it and
// formats it as a Status struct type
func StatusLink(URL string) (Status, error) {
	// Obtain all the child links
	// for a given URL
	links, err := database.GetChildLinks(URL)
	if err != nil {
		return Status{}, err
	}

	// Instantiate the Status struct
	// for our given links entry
	var status Status
	status.URL = URL

	// Iterate over each child link
	// and then recursively process
	// the child link's child URLs
	for _, link := range links {
		child, err := StatusLink(link)
		if err != nil {
			return Status{}, err
		}
		status.Links = append(status.Links, child)
	}

	return status, err
}
