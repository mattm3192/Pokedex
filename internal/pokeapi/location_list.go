package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mattm3192/Pokedex/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (PokeLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationsResp := PokeLocation{}

	cachedData, exists := cache.Get(url)
	if exists {
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return PokeLocation{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PokeLocation{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeLocation{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokeLocation{}, err
	}
	cache.Add(url, dat)

	return locationsResp, nil
}
