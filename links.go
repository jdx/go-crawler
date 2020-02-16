package crawler

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// getAllHrefs takes a goquery doc and returns all the anchor tag hrefs on the page
func getAllHrefs(doc *goquery.Document) []string {
	var hrefs []string
	// find all the anchor tags
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		hrefs = append(hrefs, href)
	})
	return hrefs
}

// absoluteURL gets the full URL of a partial URL from a base URL
// (e.g.: `/legal/cookie-notice` instead of `https://monzo.com/legal/cookie-notice`)
func absoluteURL(inputURL string, baseURL string) (string, error) {
	hostname, err := getHostname(inputURL)
	if err != nil {
		return "", err
	}
	if hostname != "" {
		// input already has a host, just return the original
		return inputURL, nil
	}
	u, err := url.Parse(baseURL) // parse the baseURL which will be our return value
	if err != nil {
		return "", err // should never be hit because the baseURL would already have been parsed
	}
	if strings.HasPrefix(inputURL, "/") {
		u.Path = inputURL
	} else {
		// handle relative path like `bar/qux` or `../bar/qux` from `https://monzo.com/foo`
		u.Path = path.Join(u.Path, inputURL)
	}
	return u.String(), nil
}

// getHostname returns the hostname of a url
func getHostname(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("getHostname error inputURL: %s %w", inputURL, err)
	}
	return u.Hostname(), nil
}

// returns true if href is a different host than baseURL
func isExternalURL(href string, baseURL string) (bool, error) {
	hostname, err := getHostname(baseURL)
	if err != nil {
		return false, err
	}
	h, err := getHostname(href)
	if err != nil {
		return false, err
	}
	return h != hostname, nil
}
