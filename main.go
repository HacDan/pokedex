package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	api "github.com/hacdan/pokedex/internal/api"
)

type Command struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	previousUrl string
	nextUrl     string
}

func main() {
	commands := getCommands()
	reader := bufio.NewReader(os.Stdin)

	config := new(Config) // allocate memory for an empty struct
	config.nextUrl = "https://pokeapi.co/api/v2/location/"

	for {
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)

		command, exists := commands[input]
		if !exists {
			fmt.Println("Command not found!")
		}
		err = command.callback(config)

		if err != nil {
			panic(err)
		}
	}
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back 20 locations on the map",
			callback:    commandMapb,
		},
	}
}

func commandExit(config *Config) error {
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	commands := getCommands() //TODO: Switch to constant of commands at runtime?

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println() //TODO: Find a better way to do this

	return nil
}

func commandMap(config *Config) error {
	locations := api.GetLocations(config.nextUrl)
	config.nextUrl = locations.NextUrl
	config.previousUrl = locations.PreviousUrl

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

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
