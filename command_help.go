package main

import (
	"fmt"

	"github.com/mattm3192/Pokedex/internal/pokecache"
)

func commandHelp(cfg *config, cache *pokecache.Cache) error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
