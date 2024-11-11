package htmldownloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Downloader interface {
	Download()
}

type HtmlDownloader struct {
	dirPath 		string
	numDownloads 	int
}

// HtmlDownloader methods

func New() (*HtmlDownloader, error) {
	return &HtmlDownloader{
		numDownloads: 0,
	}, nil
}

func (html_downloader *HtmlDownloader) Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body: " + err.Error())
		return nil, errors.New("Error reading body: " + err.Error())
	}

	file, err := os.Create(fmt.Sprintf("html%d.txt", html_downloader.numDownloads))
	if err != nil {
		fmt.Println("Error creating new file: " + err.Error())
		return nil, errors.New("Error creating new file: " + err.Error())
	}
	defer file.Close()

	// Write the HTTP status line
	_, err = file.WriteString(fmt.Sprintf("%s %s\n", resp.Proto, resp.Status))
	if err != nil {
		fmt.Println("Error writing HTTP status line: " + err.Error())
		return nil, errors.New("Error creating new file: " + err.Error())
	}

	// Write the HTTP headers
	for key, values := range resp.Header {
		for _, value := range values {
			_, err = file.WriteString(fmt.Sprintf("%s: %s\n", key, value))
			if err != nil {
				fmt.Println("Error writing headers: " + err.Error())
				return nil, errors.New("Error creating new file: " + err.Error())
			}
		}
	}

	// Separate headers from body
	_, err = file.WriteString("\n")
	if err != nil {
		fmt.Println("Error separating headers from body: " + err.Error())
		return nil, errors.New("Error creating new file: " + err.Error())
	}

	// Write the HTTP body
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error writing body: " + err.Error())
		return nil, errors.New("Error creating new file: " + err.Error())
	}

	html_downloader.numDownloads++

	return body, nil
}