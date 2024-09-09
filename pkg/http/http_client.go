package http

import (
	"bytes"
	"io"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (c *HttpClient) Get(url string) (string, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *HttpClient) Post(url string, data []byte) (string, error) {
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
