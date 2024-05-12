package main

import (
	"fmt"
)

func commandPokedex(config *Config, args []string) error {
	for name, _ := range config.pokedex {
		fmt.Println(" - ", name)
	}
	return nil
}
