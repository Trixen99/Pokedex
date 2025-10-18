package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//type config struct {
//next     string
//previous string
//}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type LocationArea struct {
	id         int
	name       string
	game_index int
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
			callback:    commandMap,
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
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area/", nil)
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()

	LocArea := LocationArea{}
	err := json.Unmarshal(res.Body, &LocationArea)

	return nil

}

//func commandMapb() error {

//}
