package main

import (
	"time"

	"github.com/e-mar404/pokedex/internal/commands"
	"github.com/e-mar404/pokedex/internal/pokeapi"
)

func main() {
	config := &commands.Config{
		PokeClient: *pokeapi.NewClient(5 * time.Second),
		Pokedex:    pokeapi.NewPokedex(),
	}
	startRepl(config)
}
