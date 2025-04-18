package commands

import "github.com/e-mar404/pokedex/internal/pokeapi"

type Config struct {
	PokeClient pokeapi.Client
	Pokedex    *pokeapi.Pokedex
	prevURL    string
	nextURL    string
}

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}

func List() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			Callback:    exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    help,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in Pokemon",
			Callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the prev 20 location areas in Pokemon",
			Callback:    mapb,
		},
		"explore": {
			name:        "explore",
			description: "List all of the Pokemon in an area",
			Callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokeball at a pokemon to simulate a catch try",
			Callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect previously caught pokemon",
			Callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show the pokemon currently in your pokedex",
			Callback:    pokedex,
		},
	}
}
