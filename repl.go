package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"github.com/Trixen99/Pokedex/internal"
)

func startRepl() {
	commandLine := bufio.NewScanner(os.Stdin)
	initValidCommands()

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

func cleanInput(text string) []string {
	finalSlice := []string{}
	splitText := strings.Split(text, " ")
	for _, word := range splitText {
		tmpWord := strings.TrimSpace(word)
		if len(tmpWord) != 0 {
			tmpWord = strings.ToLower(tmpWord)
			finalSlice = append(finalSlice, tmpWord)
		}
	}
	return finalSlice
}
