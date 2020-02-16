package crawler

import (
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Job struct{
	url string
	depth int
}

type JobResult struct{
	*Job
	doc *goquery.Document
	err error
}

// newFetchDocumentPool creates a group of HTTP fetchers to concurrently call httpFetchDocument
func (c *crawler) startWorkers(numWorkers int) <-chan *JobResult {
	out := make(chan *JobResult)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go c.startWorker(&wg, out)
	}
	go func() {
		// close the output channel when all the workers are done
		wg.Wait()
		close(out)
	}()
	return out
}

func (c *crawler) startWorker(wg *sync.WaitGroup, out chan<- *JobResult) {
	defer wg.Done()
	for job := range c.jobs {
		doc, err := httpFetchDocument(job.url)
		retries := c.maxRetries
		for err != nil && retries > 0 {
			// keep retrying while we're getting errors and have retries
			// TODO: exponential backoff
			doc, err = httpFetchDocument(job.url)
		}
		if err != nil {
			out <- &JobResult{Job: job, err: err}
			continue
		}
		out <- &JobResult{Job: job, doc: doc}
	}
}
