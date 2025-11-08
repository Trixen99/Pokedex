package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var CLText []string

func startRepl() {
	commandLine := bufio.NewScanner(os.Stdin)
	initValidCommands()

	for {
		fmt.Print("Pokedex > ")
		commandLine.Scan()
		CLText = cleanInput(commandLine.Text())
		command, ok := validCommands[CLText[0]]
		if !ok {
			fmt.Printf("Command: '%v' not recognised\n", CLText[0])
		} else {
			callCommand(CLText, command)
			//err := command.callback()
			//if err != nil {
			//	fmt.Println(err)
			//}
		}
	}
}

func callCommand(text []string, command cliCommand) {
	switch text[0] {
	case "explore":
		if len(text) <= 1 {
			fmt.Println("not enough arguments provided for the command")
			return
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
			return
		}

	default:
		err := command.callback()
		if err != nil {
			fmt.Println(err)
			return
		}
		return

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
