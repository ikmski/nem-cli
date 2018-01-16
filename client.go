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
	"path"
)

type client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

func newClient(urlStr string, httpClient *http.Client) (*client, error) {

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	client := &client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}

	return client, nil
}

func newDefaultClient() (*client, error) {

	return newClient(fmt.Sprintf("http://%s:7890", nisAddress), http.DefaultClient)
}

func (c *client) newRequest(method string, spath string, body io.Reader) (*http.Request, error) {

	url := *c.URL
	url.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	req = req.WithContext(ctx)

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

	req, err := c.newRequest(method, spath, nil)
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
