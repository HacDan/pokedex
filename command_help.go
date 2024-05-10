package main

import "fmt"

func commandHelp(config *Config, args []string) error {
	commands := getCommands() //TODO: Switch to constant of commands at runtime?

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println() //TODO: Find a better way to do this

	return nil
}
