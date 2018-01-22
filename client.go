package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type client struct {
	Host       string
	Port       string
	HTTPClient *http.Client
}

func newClient(host string, port string, httpClient *http.Client) (*client, error) {

	client := &client{
		Host:       host,
		Port:       port,
		HTTPClient: httpClient,
	}

	return client, nil
}

func newDefaultClient() (*client, error) {

	return newClient(nisAddress, "7890", http.DefaultClient)
}

func (c *client) newRequest(method string, spath string, body io.Reader) (*http.Request, error) {

	urlStr := fmt.Sprintf("http://%s:%s/%s", c.Host, c.Port, spath)

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, parsedURL.String(), body)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func outputBody(res *http.Response) error {

	defer res.Body.Close()
	inputBuf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var outputBuf bytes.Buffer
	err = json.Indent(&outputBuf, inputBuf, "", "    ")
	if err != nil {
		return err
	}

	fmt.Print(outputBuf.String())

	return nil
}

func (c *client) sendRequest(method string, spath string, body io.Reader) error {

	req, err := c.newRequest(method, spath, body)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	err = outputBody(res)
	if err != nil {
		return err
	}

	return nil
}
