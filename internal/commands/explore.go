package commands

import "fmt"

func explore(config *Config, params []string) error {
	if len(params) < 1 {
		return fmt.Errorf("please provide an area to explore")
	}
	if len(params) > 1 {
		return fmt.Errorf("please provide only 1 area to explore")
	}

	fmt.Printf("Exploring %s...\n", params[0])

	pokemonList, err := config.PokeClient.PokemonAtLocation(params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")
	for _, pokemon := range pokemonList.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
