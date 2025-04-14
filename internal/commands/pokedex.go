package commands

import "fmt"

func pokedex(config *Config, _ []string) error {
	fmt.Println("Pokemon in your pokedex:")
	fmt.Printf("count: %d\n", len(config.Pokedex.Pokemon))
	fmt.Println("=======================================")
	for _, pokemon := range config.Pokedex.Pokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
