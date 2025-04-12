package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt()

	for scanner.Scan() {
		raw := scanner.Text()
		clean := cleanInput(raw)
		inputCommand := clean[0]

		command, ok := commands()[inputCommand]
		if !ok {
			fmt.Printf("Unkown command\n")
			break
		}

		err := command.callback()
		if err != nil {
			fmt.Println(err)
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
