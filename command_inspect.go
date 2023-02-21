package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		return errors.New("you haven't caught that pokemon yet")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokeType.Type.Name)
	}
	return nil
}
