package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commandLine := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		commandLine.Scan()
		text := cleanInput(commandLine.Text())
		//fmt.Printf("Your command was: %v\n", text[0])
		command, ok := validCommands[text[0]]
		if !ok {
			fmt.Printf("Command: '%v' not recognised\n", text[0])
		} else {
			command.callback()
		}
	}
}
