package main

import "fmt"



func commandPokedex(cfg *config) error {
	fmt.Print("your pokedex contains:\n")
	for key, _ := range cfg.pokemon {
		fmt.Printf("  -%s\n", key)
	}
	return nil
}
