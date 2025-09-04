package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you didn't provide a name of pokemon to catch")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemonDetails, err := cfg.pokeapiClient.PokemonGet(name)
	if err != nil {
		return err
	}
	baseExp := pokemonDetails.BaseExperience
	if baseExp < 40 {
		baseExp = 40
	}
	catchChance := rand.Intn(baseExp)
	if catchChance > 40 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[name] = pokemonDetails
	}
	return nil
}
