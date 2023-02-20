package main

import (
	"time"

	"github.com/hereswilson/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	caughtPokemon := map[string]pokeapi.Pokemon{}
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: caughtPokemon,
	}

	startRepl(cfg)
}
