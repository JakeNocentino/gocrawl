package parser

import (
	"bytes"
	"fmt"
	"errors"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
)

// Define a ContentType Enum
type ContentType int

const (
	HTML ContentType = iota
)

var contentType = map[ContentType]string {
	HTML: "html",
}

func (content_type ContentType) String() string {
	return contentType[content_type]
}

// Interface to define a general parser
type Parser interface {
	Parse(resp http.Response) ([]string, error)
}

type HtmlParser struct {
	element string
	key string
}

// HtmlParser Methods

func New(element string, key string) (*HtmlParser, error) {
	return &HtmlParser{
		element: element,
		key: key,
	}, nil
}

func (htmlParser *HtmlParser) Parse(content []byte, baseUrl string) ([]string, error) {
	links := []string{}
	reader := bytes.NewReader(content)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			if tokenizer.Err().Error() != "EOF" {
				fmt.Printf("Error parsing content: %s\n", tokenizer.Err().Error())
				return nil, errors.New("Error parsing content: " + tokenizer.Err().Error())
			}
			return links, nil
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == htmlParser.element {
				for _, attr := range token.Attr {
					if attr.Key == htmlParser.key {
						link := attr.Val
						// Check if valid URL
						if u, err := url.Parse(link); err == nil {
							// Resolve relative URLs
							if !u.IsAbs() {
								u, _ = url.Parse(baseUrl)
								link = u.ResolveReference(u).String()
							}
						}
						links = append(links, link)
					}
				}
			}
		}
	}
}