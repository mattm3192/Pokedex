package main

import (
	"time"

	"github.com/mattm3192/Pokedex/internal/pokeapi"
	"github.com/mattm3192/Pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	newcache := pokecache.NewCache(60 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg, newcache)
}
