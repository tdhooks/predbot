package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var ErrNotFound = fmt.Errorf("not found")

type OmedaClient struct {
	client  *http.Client
	baseURL string
}

func NewOmedaClient(baseURL string, timeout time.Duration) *OmedaClient {
	return &OmedaClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

func (o *OmedaClient) Get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", o.baseURL, path)

	resp, err := o.client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return []byte{}, ErrNotFound
		} else {
			return []byte{}, fmt.Errorf("got status %v", resp.StatusCode)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("reading response body: %w", err)
	}

	return body, nil
}
