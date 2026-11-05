package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"poke2/internal/pokecache"
)


type config struct {
	commands map[string]cliCommand
	next string
	previous string
	cache pokecache.Cache
	fullCom []string
	pokemon map[string]Pokemon
}


func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			cfg.fullCom = words
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}


func getCommands() map[string]cliCommand{
	return map[string]cliCommand {
		"map": {
			name: "map",
			description: "Lists next 20 poke locations",
			callback: commandMap,
		},
		"mapb": {
			name: "map",
			description: "Lists previous 20 poke locations",
			callback: commandMapb,
		},
		"catch": {
			name: "catch",
			description: "Attempt to capture a named pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a captured pokemon for its stats",
			callback: commandInspect,
		},
		"explore": {
			name: "explore",
			description: "Finds pokemon within found location (seperate words with '-')",
			callback: commandExplore,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all the pokemon you have captured",
			callback: commandPokedex,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		
	}
}
