package main

import (
	"fmt"
)

func commandInspect(config *Config, args []string) error {

	pokemon, exists := config.pokedex[args[0]]
	if !exists {
		fmt.Println("You have not caught that", args[0], "yet!")
		return nil
	}

	// Name: pidgey
	// Height: 3
	// Weight: 18
	// Stats:
	//   -hp: 40
	//   -attack: 45
	//   -defense: 40
	//   -special-attack: 35
	//   -special-defense: 35
	//   -speed: 56
	// Types:
	//   - normal
	//   - flying
	//

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Println("  - ", stat.Stat.Name, ":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Println("  - ", pType.Type.Name)
	}
	return nil
}
