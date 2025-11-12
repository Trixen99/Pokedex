package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Trixen99/Pokedex/internal"
)

var GETCache *internal.Cache = internal.NewCache(time.Second * 60)

var validCommands = map[string]cliCommand{}

func initValidCommands() {
	validCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of all Pokemon found in the requested location in the Pokemon World",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempts to catch the designated Pokemon",
			callback:    commandCatch,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Status code: 0")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range validCommands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap() error {
	var url string
	if mapCFG.next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = mapCFG.next
	}

	var mapbyte []byte

	mapbyte, ok := GETCache.Get(url)
	if !ok {
		locationData, err := httpClientRequest("GET", url)
		//getLocationData(url)
		if err != nil {
			return err
		}
		mapbyte = locationData
	}

	mapList, err := convertLocationData(mapbyte)
	if err != nil {
		return fmt.Errorf("error received: %v", err)
	}

	if mapList.Previous != nil {
		mapCFG.previous = *mapList.Previous
	} else {
		mapCFG.previous = ""
	}

	if mapList.Next != nil {
		mapCFG.next = *mapList.Next
	} else {
		mapCFG.next = ""
	}

	for _, loc := range mapList.Results {
		fmt.Println(loc.Name)
	}

	fmt.Print("\n")
	return nil

}

func commandMapb() error {
	if mapCFG.previous == "" {
		fmt.Println("you’re on the first page")
		return nil
	}
	url := mapCFG.previous

	var mapbyte []byte

	mapbyte, ok := GETCache.Get(url)
	if !ok {
		locationData, err := httpClientRequest("GET", url)
		//getLocationData(url)
		if err != nil {
			return err
		}
		mapbyte = locationData
	}

	mapList, err := convertLocationData(mapbyte)
	if err != nil {
		return fmt.Errorf("error received: %v", err)
	}

	if mapList.Previous != nil {
		mapCFG.previous = *mapList.Previous
	} else {
		mapCFG.previous = ""
	}

	if mapList.Next != nil {
		mapCFG.next = *mapList.Next
	} else {
		mapCFG.next = ""
	}

	for _, loc := range mapList.Results {
		fmt.Println(loc.Name)
	}
	fmt.Print("\n")
	return nil

}

func convertLocationData(body []byte) (Named, error) {
	var mapList Named
	if err := json.Unmarshal(body, &mapList); err != nil {
		var errNamed Named
		return errNamed, fmt.Errorf("error: %v", err)
	}

	return mapList, nil
}

func commandExplore() error {
	location := CLText[1]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", location)
	body, err := httpClientRequest("GET", url)
	if err != nil {
		return fmt.Errorf("error with 'GET' request: %v", err)
	}

	var currentLocation LocationArea
	if err := json.Unmarshal(body, &currentLocation); err != nil {
		return fmt.Errorf("error: %v", err)
	}
	for i, _ := range currentLocation.Pokemon_encounters {
		fmt.Printf("- %v\n", currentLocation.Pokemon_encounters[i].Pokemon.Name)
	}
	fmt.Print("\n")

	return nil
}

func commandCatch() error {
	pokemon := CLText[1]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemon)
	data, err := httpClientRequest("GET", url)
	if err != nil {
		return err
	}
	var pokemonData Pokemon
	if err := json.Unmarshal(data, &pokemonData); err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %v", pokemonData.Name)
	time.Sleep(time.Second * 1)
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(time.Second * 1)
	}
	fmt.Print("\n")

	base := pokemonData.Base_experience
	if rand.Intn(base) >= base-((base/100)*45) {
		fmt.Printf("%v was caught!\n", pokemonData.Name)
	} else {
		fmt.Printf("%v escaped!\n", pokemonData.Name)
	}

	return nil
}
