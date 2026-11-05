package main

import (
	"net/http"
	"encoding/json"
	"log"
	"io"
	"fmt"
	"errors"
)


type Maps struct {
	Count    int    
	Next     string 
	Previous string 
	Results  []struct {
		Name string 
		URL  string 	
	} 
}



func commandMap(cfg *config) error {
	body, exists := cfg.cache.Get(cfg.next)
	if exists != true{

		res, err := http.Get(cfg.next)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		cfg.cache.Add(cfg.next, body)

	}
	maps := Maps{}
	err := json.Unmarshal(body, &maps)
	if err != nil {
		fmt.Println(err)
	}

	cfg.previous = cfg.next
	cfg.next = maps.Next

	for _, item := range maps.Results{
	fmt.Printf("%s\n", item.Name)
	}
	return nil
}



func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		return errors.New("previous locations not set")
	}
	body, exists := cfg.cache.Get(cfg.previous)

	if exists != true {
		res, err := http.Get(cfg.previous)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		cfg.cache.Add(cfg.previous, body)
	}
	
	maps := Maps{}
	err := json.Unmarshal(body, &maps)
	if err != nil {
		fmt.Println(err)
	}

	cfg.next = cfg.previous
	cfg.previous = ""

	for _, item := range maps.Results{
	fmt.Printf("%s\n", item.Name)
	}
	return nil
}


