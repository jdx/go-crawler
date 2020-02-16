package crawler

// Result is the element returned during crawls
// It will either have a URL or Err but not both
type Result struct{
	URL string
	Err error
}

// URLResult creates a new URL result
func URLResult(url string) *Result {
	return &Result{URL: url}
}

// ErrResult creates a new error result
func ErrResult(err error) *Result {
	return &Result{Err: err}
}
