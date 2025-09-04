package main

import "fmt"

func commandExplore(cfg *config, parameter string) error {
	if parameter == "" {
		return fmt.Errorf("you didn't provide a location to explore, use map command to see available options")
	}

	fmt.Printf("Exploring %s...\n", parameter)
	locationDetails, err := cfg.pokeapiClient.ListLocationDetails(&parameter)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
