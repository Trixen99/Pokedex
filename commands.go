package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Trixen99/Pokedex/internal"
)

var mapCache *internal.Cache = internal.NewCache(time.Second * 60)

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
			name:        "Explore",
			description: "Displays the names of all Pokemon found in the requested location in the Pokemon World",
			callback:    commandExplore,
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

	mapbyte, ok := mapCache.Get(url)
	if !ok {
		locationData, err := getLocationData(url)
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

	mapbyte, ok := mapCache.Get(url)
	if !ok {
		locationData, err := getLocationData(url)
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

func getLocationData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error with 'GET' request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	mapCache.Add(url, body)

	return body, nil
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
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/pastoria-city-area")
	if err != nil {
		return fmt.Errorf("error with 'GET' request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error with 'GET' request: %v", err)
	}
	var currentLocation LocationArea
	if err := json.Unmarshal(body, &currentLocation); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println(currentLocation)
	return nil
}
