package parser

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetHyperlinks process a given URL and the request to the UR
// to obtain all the links in it. It does so by tokenizing the 
// result <html> body and identifies hyperlinks in it
func GetHyperlinks(BaseURL string, response *http.Response) []string {
	var links []string
	token := html.NewTokenizer(response.Body)
	for {
		tokenIterator := token.Next()
		tokenTest := token.Token()
		switch {
		case tokenIterator == html.ErrorToken:
			return links
		case tokenIterator == html.StartTagToken:
			if tokenTest.Data != "a" {
				continue
			}

			url := GetLinkHelper(tokenTest.Attr)

			// If parsed url is a redundant 
			// or the same link discard it
			if url == "" || url == "/" || url == "#" || url == BaseURL {
				continue
			}

			// Filter only for http links
			if strings.Index(url, "http") == 0 {
				links = append(links, url)
				continue
			}

			// In the case of relative links
			// prepend the original URL and 
			// check if its valid
			if ValidateURL(BaseURL+url) == nil {
				links = append(links, BaseURL+url)
			}
		}
	}
}

// GetLinkHelper iterates over the token attributes 
// to search for the link type token
func GetLinkHelper(tokenAttribute []html.Attribute) string {
	if len(tokenAttribute) == 0 {
		return ""
	}

	if tokenAttribute[0].Key == "href" {
		return tokenAttribute[0].Val
	}

	return GetLinkHelper(tokenAttribute[1:])
}

// FilterHyperlinks filters out links that
// lead away from the original domain URL
func FilterHyperlinks(URL string, links []string) []string {
	index := 0
	for _, link := range links {
		if strings.Contains(link, URL) {
			links[index] = link
			index++
		}
	}
	return links[:index]
}
