package commands

import "fmt"

func help(_ *Config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, command := range List() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
