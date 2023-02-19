package pokeapi

import (
	"net/http"
	"time"
)

// Client is a client for the PokeAPI.
type Client struct {
	// HTTP client used to communicate with the API.
	client http.Client
}

// NewClient returns a new PokeAPI client.
func NewClient(timeout time.Duration) Client {
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
	}
}
