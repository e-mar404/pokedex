package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	
	for _, command := range commands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
