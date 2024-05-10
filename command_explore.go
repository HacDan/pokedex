package main

import (
	"fmt"
)

func commandExplore(config *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	pokemonFromLocation := config.pokeclient.GetPokemonFromLocation(url)
	fmt.Println("Exploring ", args[0], "...")
	for _, pokemon := range pokemonFromLocation.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	return nil
}
