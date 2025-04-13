package commands

import "fmt"

func catch(config *Config, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("You can only try to catch one pokemon at a time")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])

	pokemon, caught, err := config.PokeClient.CatchPokemon(params[0])
	if err != nil {
		return err
	}

	var message string
	if caught {
		message = fmt.Sprintf("%s was caught!", params[0])
		config.Pokedex[params[0]] = pokemon
	} else {
		message = fmt.Sprintf("%s escaped", params[0])
	}

	fmt.Println(message)

	return nil
}
