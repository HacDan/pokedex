package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
