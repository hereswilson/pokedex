package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// LocationList returns a list of locations.
func (c *Client) ListLocations(pageURL *string) (locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := locations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return locations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locations{}, err
	}

	locationsResp := locations{}

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return locations{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
