# GoCrawl

A simple yet extensible web crawler built in GoLang.

This project was started as a way to learn GoLang and create a simple web crawler. However, a web crawler can be as simple as a single-day endeavor and as complex as a team-wide production-esque effort. As such, iterations on this project are essentially endless, which I believe is what makes a project exciting to work on.

This project is written in GoLang. To install GoLang, visit this [webpage](https://go.dev/doc/install)

### Quickstart

Web crawler functionality currently is not yet complete. To run what exists (initializing a queue for URLs, crawling a single seed URL):

```bash
cd cmd/gocrawl
go build gocrawl.go
sudo ./gocrawl
```