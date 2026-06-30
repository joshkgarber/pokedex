package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Println("Your command was:", commandName)
	}
}

func cleanInput(text string) []string {
	loweredInputs := strings.ToLower(text)
	splitInputs := strings.Fields(loweredInputs)
	return splitInputs
}

