package data

import (
	"net/http"
	"time"
)

// TransparencyClient is used to wrap calls to relays.
type TransparencyClient struct {
	clt     http.Client
	baseURL string
}

// NewTransparencyClient creates a new TransparencyClient using a relay URL and a request timeout.
func NewTransparencyClient(baseURL string, timeout time.Duration) *TransparencyClient {
	return &TransparencyClient{
		clt: http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}
