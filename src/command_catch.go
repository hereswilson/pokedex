package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	rand_int := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at  %s...\n", pokemon.Name)
	if rand_int > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("You caught %s!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
