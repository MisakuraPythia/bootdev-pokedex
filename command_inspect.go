package main

import (
	"fmt"
)


func commandInspect(cfg *config) error {
	poke := cfg.fullCom[1]
	val, ok := cfg.pokemon[poke]
	if ok == false {
		return fmt.Errorf("you have not captured a %s", poke)
	}else {
		fmt.Printf("Name: %s\nWeight: %d\nStats:\n", val.Name, val.Weight)
		for _, stat := range val.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, poketype := range val.Types {
			fmt.Printf("  -%s\n", poketype.Type.Name)
		}

	}
	return nil
}
