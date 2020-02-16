package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"crawler"
)

var (
	concurrency = 1
	depth = -1
	retries = 0
	halt = false
)

var app = &cli.App{
	Usage:   "crawl a webpage and recursively print all the links it has on the same subdomain",
	Version: crawler.Version,
	Authors: []*cli.Author{{Name: "Jeff Dickey", Email: "jeff@dickey.us"}},
	ArgsUsage: "[URL...]",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:        "max-concurrency",
			Aliases: []string{"c"},
			Usage:       "maximum scrape `workers` to run simultaneously",
			Value:       concurrency,
			Destination: &concurrency,
		},
		&cli.IntFlag{
			Name:        "max-depth",
			Aliases: []string{"d"},
			Usage:       "maximum `depth` of page from starting URL. Set to \"-1\" for no limit.",
			Value:       depth,
			Destination: &depth,
		},
		&cli.IntFlag{
			Name:        "max-retries",
			Aliases: []string{"r"},
			Usage:       "maximum `times` to retry a given URL before displaying an error",
			Value:       retries,
			Destination: &retries,
		},
		&cli.BoolFlag{
			Name: "halt",
			Usage: "stop on first error (not included successful retries)",
			Destination: &halt,
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			// didn't specify any URLs to scrape so show usage
			cli.ShowAppHelpAndExit(c, 1)
		}
		urls := c.Args().Slice()

		// start crawling
		links := crawler.Crawl(
			urls,
			crawler.MaxConcurrency(concurrency),
			crawler.MaxDepth(depth),
			crawler.MaxRetries(retries),
		)

		// output links as we receive them
		// will close automatically unless --max-depth is `-1`
		for res := range links {
			if res.Err != nil {
				if halt {
					return res.Err
				}
				log.Println(res.Err) // emit error to stderr
				continue
			}
			fmt.Println(res.URL)
		}
		return nil
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
