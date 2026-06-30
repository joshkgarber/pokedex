package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// Examples:
	// "hello world" -> ["hello", "world"]
	// "Charmander Bulbasur PIKACHU" -> ["charmander", "bulbasur", "pikachu"]

	loweredInputs := strings.ToLower(text)
	splitInputs := strings.Fields(loweredInputs)

	return splitInputs
}

