package main

import (
	"fmt"
	"os"
)

type commandStart interface {
	startcommand()
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var validCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandExit,
	},
}

//func (c cliCommand) startCommand {

//}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Status code: 0")
}

/*func commandHelp() error {
	fmt.Println("Usage:")
	for _, command := range validCommands {
		fmt.Printf("%v: %v", command.name, command.description)
	}
	return nil
}
*/
