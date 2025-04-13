package commands

import (
	"fmt"
	"os"
)

func exit(_ *Config, _ []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
