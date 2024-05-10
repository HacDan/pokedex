package main

import (
	"fmt"
)

func commandMap(config *Config, args string) error {
	locations := config.pokeclient.GetLocations(config.nextUrl)

	config.nextUrl = locations.NextUrl
	config.previousUrl = locations.PreviousUrl

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
