package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt()

	for scanner.Scan(){
		raw := scanner.Text()
		clean := cleanInput(raw)
		fmt.Printf("Your command was: %s\n", clean[0])

		prompt()
	}
}

func prompt() {
	fmt.Print("Pokedex > ")
}
