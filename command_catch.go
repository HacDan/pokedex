package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	pokemonFromLocation := config.pokeclient.GetPokemon(url)
	fmt.Println("Throwing a Pokeball at ", args[0], "...")

	chance := rand.Intn(pokemonFromLocation.BaseExperience)

	if chance > 40 {
		fmt.Println(pokemonFromLocation.Name, " escaped!")
		return nil
	}

	config.pokedex[pokemonFromLocation.Name] = pokemonFromLocation
	fmt.Println(pokemonFromLocation.Name, " was caught!")
	return nil
}
