// Package goanda implements a stateless client for the oanda REST-v20 API.
package goanda

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Client is a stateless client for the oanda REST-v20 API.
type Client struct {
	rc *resty.Client
}

// New returns a new client configured with the necessary auth and address etc.
func NewClient(key string) (c *Client) {
	rc := resty.New().
		SetHostURL(HostURL).
		SetAuthToken(key)

	return &Client{rc: rc}
}

// GetHTTPClient returns the underlying http client used by the resty client.
func (c *Client) GetHTTPClient() *http.Client {
	return c.rc.GetClient()
}
