package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you didn't provide a name of pokemon to inspect")
	}
	name := args[0]

	_, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon yet")
	}
	fmt.Printf("Name: %s\n", cfg.caughtPokemon[name].Name)
	fmt.Printf("Height: %d\n", cfg.caughtPokemon[name].Height)
	fmt.Printf("Weight: %d\n", cfg.caughtPokemon[name].Weight)

	fmt.Println("Stats:")
	for _, stat := range cfg.caughtPokemon[name].Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketype := range cfg.caughtPokemon[name].Types {
		fmt.Printf("  - %s\n", poketype.Type.Name)
	}
	return nil
}
