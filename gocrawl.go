package gocrawl

import (
	"github.com/jakenocentino/gocrawl/urlfrontier"
	"github.com/jakenocentino/gocrawl/htmldownloader"
	"github.com/jakenocentino/gocrawl/parser"

)

type GoCrawler struct {
	downloader 		*htmldownloader.HtmlDownloader
	parser 			*parser.HtmlParser
	urlFrontier 	*urlfrontier.UrlFrontier
	depth 			int
}

func New(depth int) (*GoCrawler, error) {
	urlFrontier, _ := urlfrontier.New()
	htmlDownloader, _ := htmldownloader.New()
	parser, _ := parser.New("a", "href")
	return &GoCrawler{
		downloader: htmlDownloader,
		parser: parser,
		urlFrontier: urlFrontier,
		depth: depth,
	}, nil
}

func (crawler *GoCrawler) Crawl(seedUrl string) error {
	// Initiate crawl with provided seed URL
	crawler.urlFrontier.AddUrl(seedUrl)
	curDepth := 0
	for !crawler.urlFrontier.IsEmpty() && curDepth < crawler.depth {
		size, _ := crawler.urlFrontier.Size()
		for i := 0; i < size; i++ {
			// Get latest URL in URL Frontier Queue
			url, _ := crawler.urlFrontier.PopUrl()

			// Download the content and return the content for the parser to use
			body, _ := crawler.downloader.Download(url)

			// Parse the content for additional URLs to crawl
			links, _ := crawler.parser.Parse(body, url)

			// Add the respective links to the URL Frontier
			crawler.urlFrontier.AddAllUrls(links)
		}
		// Add one to the current depth traversed
		curDepth++
	}

	return nil
}