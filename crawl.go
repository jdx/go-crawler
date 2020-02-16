package crawler

import (
	"sync"
)

// Crawl startURLs and emit results for every link that is found recursively
func Crawl(startURLs []string, options ...Option) <-chan *Result {
	c := newCrawler(options)

	// populate with start URLs
	for _, startURL := range startURLs {
		c.addJob(startURL, 0)
	}

	docChan := c.startWorkers(c.maxConcurrency)
	go c.processDocs(docChan)
	go c.closeWhenDone()

	return c.results
}

// stores some state for a crawl run
type crawler struct{
	// options
	maxConcurrency int
	maxDepth int
	maxRetries int

	jobs chan *Job
	results chan *Result
	seen set
	wg sync.WaitGroup
}

// creates a crawler struct from option funcs
func newCrawler(opts []Option) *crawler {
	c := &crawler{
		maxConcurrency: 1,
		maxDepth: -1,
		maxRetries: 0,
		jobs: make(chan *Job),
		results: make(chan *Result),
		seen: set{},
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.maxConcurrency <= 0 {
		c.maxConcurrency = 1
	}
	if c.maxRetries < 0 {
		c.maxRetries = 0
	}
	return c
}

// add a new URL to be fetched
func (c *crawler) addJob(url string, depth int) {
	if c.maxDepth != -1 && depth > c.maxDepth {
		return
	}
	c.wg.Add(1)
	go func() {
		// done async because backpressure here may cause deadlocking
		// waiting for jobs to complete but also for the job to be added
		c.jobs <- &Job{url: url, depth: depth}
	}()
}

// processDocs takes a goquery document from a job and pulls the hrefs out then processes them
func (c *crawler) processDocs(in <-chan *JobResult) {
	defer close(c.results)
	for jr := range in {
		if jr.err != nil {
			c.emitError(jr.err)
		} else {
			for _, href := range getAllHrefs(jr.doc) {
				c.processURL(href, jr)
			}
		}
		c.wg.Done()
	}
}

// given a URL, make sure it's valid (not external, not previously seen)
// if valid, emits it in response and adds a job to process the new URL
func (c *crawler) processURL(href string, jr *JobResult) {
	// get the full URL
	href, err := absoluteURL(href, jr.url)
	if err != nil {
		c.emitError(err)
		return
	}

	// filter external URLs on different domains
	isExternal, err := isExternalURL(href, jr.url)
	if err != nil {
		c.emitError(err)
		return
	}
	if isExternal {
		return
	}

	// filter out previously seen URLs
	if c.seen.Has(href) {
		return
	}
	c.seen.Add(href)
	c.emitURL(href)

	// start processing this URL
	c.addJob(href, jr.depth + 1)
}

// send a URL to the calling code via the response chan
func (c *crawler) emitURL(url string) {
	c.results <- URLResult(url)
}

// send an error to the calling code via the response chan
func (c *crawler) emitError(err error) {
	c.results <- ErrResult(err)
}

// closes the response channel when there are no more jobs
func (c *crawler) closeWhenDone() {
	c.wg.Wait()
	close(c.jobs)
}
