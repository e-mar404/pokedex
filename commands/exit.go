package commands

import (
	"fmt"
	"os"
)

func commandExit(_ *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
