package main

import "fmt"

func commandCatch(config *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	pokemonFromLocation := config.pokeclient.GetPokemon(url)
	fmt.Println("Throwing a Pokeball at ", args[0], "...")
	fmt.Println("Base experience: ", pokemonFromLocation.BaseExperience)

	return nil
}
