package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

type pokeLocation struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func LocationsCall(url *string) (pokeLocation, error) {
	resp, err := http.Get(*url)
	if err != nil {
		return pokeLocation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokeLocation{}, err
	}

	var locations pokeLocation
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return pokeLocation{}, err
	}

	return locations, nil
}
