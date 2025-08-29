package main

import (
	"fmt"

	"github.com/mattm3192/Pokedex/internal/pokeapi"
)

func commandMapB(cfg *config) error {
	if cfg.previousURL == "" {
		fmt.Println("You haven't used map yet")
		return nil
	} else if cfg.previousURL == "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0" {
		fmt.Println("You are on the first page")
		return nil
	}
	locations, err := pokeapi.LocationsCall(&cfg.previousURL)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	if locations.Previous != nil {
		cfg.previousURL = *locations.Previous
	} else {
		cfg.previousURL = cfg.nextURL
	}
	cfg.nextURL = locations.Next
	return nil
}
