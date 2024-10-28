package main

import (
	"flag"
	"github.com/jakenocentino/gocrawl"
)

func main() {
	// Parse an initial seed url to commence crawling
	seedUrl := flag.String("seed-url", "https://www.nocentino.dev", "The seed url used to instantiate the web crawling process")
	flag.Parse()

	crawler, _ := gocrawl.New(3)
	crawler.Crawl(*seedUrl)
}