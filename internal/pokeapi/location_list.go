package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationsResp := PokeLocation{}

	cachedData, exists := c.cache.Get(url)
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
	c.cache.Add(url, dat)

	return locationsResp, nil
}

func (c *Client) ListLocationDetails(id *string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + *id

	pokemonResp := LocationDetails{}

	cachedData, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(cachedData, &pokemonResp)
		if err != nil {
			return LocationDetails{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationDetails{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return LocationDetails{}, err
	}
	c.cache.Add(url, dat)
	return pokemonResp, nil
}
