package comms

import (
	"net/http"
	"time"
)

const httpTimeout = time.Second * 10
const maxConn = 100

// HTTPComms represents an HTTP Client Comms
type HTTPComms struct {
	client *http.Client
}

// NewHTTPComms initialize an HTTP client
func NewHTTPComms() *HTTPComms {
	netClient := &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        maxConn,
			MaxIdleConnsPerHost: maxConn,
		},
	}

	comms := &HTTPComms{
		client: netClient,
	}
	return comms
}

// Send an http request
func (c *HTTPComms) Send(r *http.Request) (*http.Response, error) {
	return c.client.Do(r)
}
