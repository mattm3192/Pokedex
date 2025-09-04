package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mattm3192/Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		commands := getCommands()
		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world, each subsequent call gets the next 20.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations areas if map has been used more than once, otherwise will tell you your on the first page. It's a way to go back.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes the name of a location area as an argument, and will show the user a list of all the Pokemon located there.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes a pokemon name as an argument and adds it to the users pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes the name of a caught pokemon as an argument and returns that pokemons stats. ",
			callback:    commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	loweredString := strings.ToLower(text)
	words := strings.Fields(loweredString)
	return words
}
