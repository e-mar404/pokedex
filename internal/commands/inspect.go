package commands

import (
	"fmt"

	"github.com/e-mar404/pokedex/internal/pokeapi"
)

func inspect(config *Config, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("you should only inspect one Pokemon at a time")
	}

	pokemon, ok := config.Pokedex.Get(params[0])
	if !ok {
		return fmt.Errorf("pokemon has not been caught before")
	}

	printDetails(pokemon)

	return nil
}

func printDetails(p pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
}
