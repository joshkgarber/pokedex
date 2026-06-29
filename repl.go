package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// Examples:
	// "hello world" -> ["hello", "world"]
	// "Charmander Bulbasur PIKACHU" -> ["charmander", "bulbasur", "pikachu"]
	var cleanInputs []string

	loweredInputs := strings.ToLower(text)
	splitInputs := strings.Split(loweredInputs, " ")

	for _, input := range splitInputs {
		trimmedInput := strings.Trim(input, " ")
		if input != "" {
			cleanInputs = append(cleanInputs, trimmedInput)
		}
	}



	return cleanInputs
}

