# GoCrawl

A simple yet extensible web crawler built in GoLang.

This project was started as a way to learn GoLang and create a simple web crawler. However, a web crawler can be as simple as a single-day endeavor and as complex as a team-wide production-esque effort. As such, iterations on this project are essentially endless, which I believe is what makes a project exciting to work on.

This project is written in GoLang. To install GoLang, visit this [webpage](https://go.dev/doc/install)

### Quickstart

To run the current main program, run:

```bash
cd cmd/gocrawl
go build gocrawl.go
./gocrawl
```

Future README improvements will describe how to use the library itself within your Go code.

### Improvements
Below is a list of future improvements, loosely ordered from highest priority to lowest priority:

1. Add delay between requests to become a more polite web crawler (UrlFrontier component).
2. Parallelize the web crawler using goroutines.
3. Find a better way to save and access the crawled HTML pages rather than saving as a local file.
4. Add UrlFilter component to define more ways to filter out URLs. Move the `visited` map here from UrlFrontier when complete.
5. Add prioritization into how we select URLs to crawl (UrlFrontier component).
6. Add Robots.txt support (HtmlDownloader component).
7. Better error handling and logging.
8. Add support for benchmarking and metrics collection.
9. Make user experience more robust.
10. Add check for content seen before URL seen, since sometimes content between different URLs can be the same.
11. Add caching component.
