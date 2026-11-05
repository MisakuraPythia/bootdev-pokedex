package main

import(
	"os"
	"fmt"
)

func commandExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
