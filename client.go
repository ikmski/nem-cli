package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

func NewClient(urlStr string, httpClient *http.Client, logger *log.Logger) (*Client, error) {

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	client := &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
		Logger:     logger,
	}

	return client, nil
}

func NewDefaultClient() (*Client, error) {

	var buf bytes.Buffer
	logger := log.New(&buf, "INFO: ", log.Lshortfile)

	return NewClient("http://185.117.22.111:7890/", http.DefaultClient, logger)
}

func (c *Client) Status() error {

	url := *c.URL
	url.Path = path.Join(c.URL.Path, "status")

	res, err := c.HTTPClient.Get(url.String())
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", res)

	return nil
}
