package main

import (
	"errors"
	"fmt"

	"github.com/mattm3192/Pokedex/internal/pokecache"
)

func commandMap(cfg *config, cache *pokecache.Cache) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextURL, cache)
	if err != nil {
		return err
	}

	cfg.nextURL = &locationsResp.Next
	cfg.previousURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache) error {
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousURL, cache)
	if err != nil {
		return err
	}

	cfg.nextURL = &locationResp.Next
	cfg.previousURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
