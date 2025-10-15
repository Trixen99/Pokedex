package main

import (
	"strings"
)

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
