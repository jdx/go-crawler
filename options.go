package crawler

// Option is a type of function that changes the behavior of a crawl
// This uses the functional options pattern:
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type Option func(*crawler)

// MaxConcurrency of simultaneous requests that can be made
func MaxConcurrency(maxConcurrency int) Option {
	return func(c *crawler) {
		c.maxConcurrency = maxConcurrency
	}
}

// MaxDepth is an Option which sets the recursion limit for the crawler
func MaxDepth(maxDepth int) Option {
	return func(c *crawler) {
		c.maxDepth = maxDepth
	}
}

// MaxRetries is an Option which sets how many times a given URL should be retried on error
func MaxRetries(maxRetries int) Option {
	return func(c *crawler) {
		c.maxRetries = maxRetries
	}
}
