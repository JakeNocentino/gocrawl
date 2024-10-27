package main

import(
	"flag"
	"github.com/jakenocentino/gocrawl/url_frontier"
	"github.com/jakenocentino/gocrawl/html_downloader"
)

func main() {
	// Parse an initial seed url to commence crawling
	seedUrl := flag.String("seed-url", "https://www.nocentino.dev", "The seed url used to instantiate the web crawling process")
	flag.Parse()

	urlfrontier, _ := url_frontier.New()
	urlfrontier.AddUrl(*seedUrl)

	htmldownloader, _ := html_downloader.New("downloads")
	htmldownloader.Download(*seedUrl)
}