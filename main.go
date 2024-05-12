package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	api "github.com/hacdan/pokedex/internal/api"
)

type Command struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	previousUrl string
	nextUrl     string
	pokeclient  api.Client
	pokedex     map[string]api.PokemonResponse
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
		"explore": {
			name:        "explore",
			description: "Explores the area around you for Pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempst to catch the pokemon named in the second argument",
			callback:    commandCatch,
		},
	}
}

func main() {
	commands := getCommands()
	reader := bufio.NewReader(os.Stdin)

	config := new(Config) // allocate memory for an empty struct
	config.nextUrl = "https://pokeapi.co/api/v2/location-area/"
	config.pokeclient = api.NewClient(10*time.Second, 5*time.Minute)
	config.pokedex = make(map[string]api.PokemonResponse)

	for {
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		cmd, args := parseCommand(input)

		command, exists := commands[cmd]
		if !exists {
			fmt.Println("Bad command or filename")
		} else {
			err = command.callback(config, args)

			if err != nil {
				panic(err)
			}
		}
	}
}

func parseCommand(s string) (string, []string) {
	s = strings.TrimSpace(s)

	splitString := strings.Split(s, " ")

	if len(splitString) == 1 {
		return splitString[0], []string{""}
	}
	return splitString[0], splitString[1:]
}
