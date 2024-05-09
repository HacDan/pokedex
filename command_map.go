package main

import (
	"fmt"

	api "github.com/hacdan/pokedex/internal/api"
)

func commandMap(config *Config) error {
	locations := api.GetLocations(config.nextUrl)
	config.nextUrl = locations.NextUrl
	config.previousUrl = locations.PreviousUrl

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
