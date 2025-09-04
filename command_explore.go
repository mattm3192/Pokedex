package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you didn't provide a location to explore, use map command to see available options")
	}

	name := args[0]
	locationDetails, err := cfg.pokeapiClient.ListLocationDetails(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
