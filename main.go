package main

import (
	"poke2/internal/pokecache"
	"time"	
)



func main() {
	cfg := &config{}

	cfg.pokemon = make(map[string]Pokemon)
	cfg.next = "https://pokeapi.co/api/v2/location-area/"
	cfg.cache = pokecache.NewCache(5 * time.Minute)
	
	startRepl(cfg)
}
