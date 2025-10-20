package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Named struct {
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
			name:        "map",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
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

	mapList, err := getLocationData(url)
	if err != nil {
		return err
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

	mapList, err := getLocationData(url)
	if err != nil {
		return err
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

func getLocationData(url string) (Named, error) {
	res, err := http.Get(url)
	if err != nil {
		var errNamed Named
		return errNamed, fmt.Errorf("error with 'GET' request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errNamed Named
		return errNamed, fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		var errNamed Named
		return errNamed, fmt.Errorf("err")
	}

	var mapList Named
	if err := json.Unmarshal(body, &mapList); err != nil {
		var errNamed Named
		return errNamed, fmt.Errorf("error3")
	}

	return mapList, nil
}
