package client

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

const (
	BaseURI = "https://www.jma.go.jp/"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		http.DefaultClient,
	}
}

func callAPI(ctx context.Context, httpClient *http.Client, path string) (*http.Response, error) {
	resp, err := ctxhttp.Get(ctx, httpClient, BaseURI+path)
	if err != nil {
		return nil, fmt.Errorf("HTTP Get error: %w", err)
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP status: %d", resp.StatusCode)
	}

	return resp, nil
}
