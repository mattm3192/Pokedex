package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokeLocation struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func LocationsCall(url *string) (PokeLocation, error) {
	resp, err := http.Get(*url)
	if err != nil {
		return PokeLocation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeLocation{}, err
	}

	var locations PokeLocation
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return PokeLocation{}, err
	}

	return locations, nil
}
