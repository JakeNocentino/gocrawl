package html_downloader

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
	dir_path 		string
	num_downloads 	int
}

// HtmlDownloader methods

func New(dir_path string) (*HtmlDownloader, error) {
	return &HtmlDownloader{
		dir_path: dir_path,
		num_downloads: 0,
	}, nil
}

func (html_downloader *HtmlDownloader) Download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = os.MkdirAll(html_downloader.dir_path, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating download directory: " + err.Error())
		return errors.New("Error creating download directory: " + err.Error())
	}

	file, err := os.Create(fmt.Sprintf("html%d.txt", html_downloader.num_downloads))
	if err != nil {
		fmt.Println("Error creating new file: " + err.Error())
		return errors.New("Error creating new file: " + err.Error())
	}
	defer file.Close()

	// Write the HTTP status line
	_, err = file.WriteString(fmt.Sprintf("%s %s\n", resp.Proto, resp.Status))
	if err != nil {
		fmt.Println("Error writing HTTP status line: " + err.Error())
		return errors.New("Error writing HTTP status line: " + err.Error())
	}

	// Write the HTTP headers
	for key, values := range resp.Header {
		for _, value := range values {
			_, err = file.WriteString(fmt.Sprintf("%s: %s\n", key, value))
			if err != nil {
				fmt.Println("Error writing headers: " + err.Error())
				return errors.New("Error writing headers: " + err.Error())
			}
		}
	}

	// Separate headers from body
	_, err = file.WriteString("\n")
	if err != nil {
		fmt.Println("Error separating headers from body: " + err.Error())
		return errors.New("Error separating headers from body: " + err.Error())
	}

	// Write the HTTP body
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing body: " + err.Error())
		return errors.New("Error writing body: " + err.Error())
	}

	fmt.Printf("Response written to %s", fmt.Sprintf("%s/html%d.txt", html_downloader.dir_path, html_downloader.num_downloads))

	return nil
}