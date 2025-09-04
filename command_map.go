package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
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

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousURL)
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
