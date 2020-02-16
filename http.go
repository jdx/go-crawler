package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// makes an HTTP GET request for a page
// then parses it with goquery
func httpFetchDocument(pageURL string) (*goquery.Document, error) {
	wrapErr := func(err error) error {
		// wrap any errors so we know which URL caused them
		return fmt.Errorf("HTTP GET[%s]: %w", pageURL, err)
	}
	res, err := http.Get(pageURL)
	if err != nil {
		return nil, wrapErr(err)
	}
	if res.StatusCode != 200 {
		return nil, wrapErr(fmt.Errorf("HTTP error %s", res.Status))
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, wrapErr(err)
	}
	return doc, nil
}

