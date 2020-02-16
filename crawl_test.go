package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrawl_Monzo(t *testing.T) {
	links := collect(Crawl([]string{"https://monzo.com"}, MaxDepth(1)))
	assert.Contains(t, links, "https://monzo.com/features/travel")
}

func collect(res <-chan *Result) []string {
	var out []string
	for r := range res {
		out = append(out, r.URL)
	}
	return out
}
