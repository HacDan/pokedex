package main

import (
	"fmt"

	api "github.com/hacdan/pokedex/internal/api"
)

func commandMapb(config *Config) error {
	if config.previousUrl == "" {
		fmt.Println("Error")
		return nil
	}

	locations := api.GetLocations(config.previousUrl)
	config.nextUrl = locations.NextUrl
	config.previousUrl = locations.PreviousUrl

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
