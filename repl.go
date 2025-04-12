package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/e-mar404/pokedex/internal/commands"
)


func startRepl(config *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	prompt()

	for scanner.Scan() {
		raw := scanner.Text()
		clean := cleanInput(raw)
		if len(clean) < 1 {
			fmt.Printf("\n")
			prompt()
			continue	
		}

		inputCommand := clean[0]

		command, ok := commands.List()[inputCommand]
		if !ok {
			fmt.Printf("Unkown command\n")
			prompt()
			continue	
		}

		err := command.Callback(config)
		if err != nil {
			fmt.Printf("command error: %v\n", err)
		}

		prompt()
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func prompt() {
	fmt.Print("Pokedex > ")
}


