# crawler

A simple web crawler that recursively fetches links

## Usage

Run this with `go run crawler/cmd/crawl [URL...]`

```
NAME:
   crawl - crawl a webpage and recursively print all the links it has on the same subdomain

USAGE:
   crawl [global options] command [command options] [URL...]

VERSION:
   1.0.0

AUTHOR:
   Jeff Dickey <jeff@dickey.us>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --max-concurrency workers, -c workers  maximum scrape workers to run simultaneously (default: 1)
   --max-depth depth, -d depth            maximum depth of page from starting URL. Set to "-1" for no limit. (default: -1)
   --max-retries times, -r times          maximum times to retry a given URL before displaying an error (default: 0)
   --halt                                 stop on first error (not included successful retries) (default: false)
   --help, -h                             show help (default: false)
   --version, -v                          print the version (default: false)
```

Example run:

```
$ go run crawler/cmd/crawl -d 0 https://google.com
https://google.com/preferences%3Fhl=en
https://google.com/search%3Fie=UTF-8&q=Susan+B.+Anthony&oi=ddle&ct=144864050&hl=en&sa=X&ved=0ahUKEwj1z-Lx5dTnAhXO854KHUy1CYcQPQgD
https://google.com/advanced_search%3Fhl=en&authuser=0
https://google.com/intl/en/ads/
https://google.com/services/
https://google.com/intl/en/about.html
https://google.com/intl/en/policies/privacy/
https://google.com/intl/en/policies/terms/
```

## Testing

Test with `go test ./...`

## Code Layout

* [`crawl.go`](crawl.go) – main entry point for crawling library
* [`http.go`](http.go) – web request logic
* [`links.go`](links.go) – helpers to filter/map URLs
* [`worker.go`](worker.go) – concurrency logic
* [`cmd/crawl/main.go`](cmd/crawl/main.go) – main entry point for CLI wrapper of crawling library
