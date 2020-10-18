package handlers

import (
	"encoding/json"
	"net/http"
)

// InputBody represents the JSON 
// POST message request body 
type InputBody struct {
	URL string `json:"URL"`
}

// GetURL extracts the "URL" 
// entry from the JSON body
func GetURL(r *http.Request) (string, error) {
	var body InputBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.URL, nil
}
