package pokeapi

import (
	"net/http"
	"time"

	"github.com/hereswilson/pokedex/internal/pokecache"
)

// Client is a client for the PokeAPI.
type Client struct {
	cache pokecache.Cache
	// HTTP client used to communicate with the API.
	client http.Client
}

// NewClient returns a new PokeAPI client.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}
