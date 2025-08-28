package main

import (
	"fmt"
	"strings"

	"github.com/mattm3192/Pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	isMapURL := false
	if cfg.nextURL == "" {
		cfg.nextURL = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	} else {

		urlParts := strings.Split(cfg.nextURL, "/")
		for _, part := range urlParts {
			if part == "location-area" {
				isMapURL = true
			}
		}
	}
	if !isMapURL {
		cfg.nextURL = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	}

	locations, err := pokeapi.LocationsCall(&cfg.nextURL)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.previousURL = cfg.nextURL
	cfg.nextURL = locations.Next
	return nil
}
