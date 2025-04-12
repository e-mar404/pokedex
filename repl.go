package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"

	"github.com/e-mar404/pokedex/commands"
)

func startRepl() {
	config := &commands.Config{}
	scanner := bufio.NewScanner(os.Stdin)
	prompt()

	for scanner.Scan() {
		raw := scanner.Text()
		clean := cleanInput(raw)
		inputCommand := clean[0]

		command, ok := commands.List()[inputCommand]
		if !ok {
			fmt.Printf("Unkown command\n")
			break
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
