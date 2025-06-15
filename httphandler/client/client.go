package client

import (
	"net/http"
	"time"
)

// APIClient struct to hold api information
type APIClient struct {
	HttpClient *http.Client
	BaseURL    string
}

// NewAPIClient returns api client struct
func NewAPIClient(baseURL string, timeout time.Duration) *APIClient {
	return &APIClient{
		HttpClient: &http.Client{Timeout: timeout},
		BaseURL:    baseURL,
	}
}
