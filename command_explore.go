package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"log"
)

func commandExplore(cfg *config) error{
	if len(cfg.fullCom) < 2 {
		return fmt.Errorf("failed to fetch location")
	}
	URL, exists := cfg.cache.Find(cfg.fullCom[1]) 
	var err error
	var res *http.Response
	if exists == true {
		res, err = http.Get(URL)
	} else {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/" + cfg.fullCom[1])
	}

	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	area := Location{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
	}

	for _, pokemon := range area.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}
	

	return nil
}


type Location struct {
	ID                   int    
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
