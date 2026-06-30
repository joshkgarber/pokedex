package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name 		string
	description	string
	callback	func() error
}

var supportedCommands map[string]cliCommand

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

		supportedCommands = make(map[string]cliCommand)

		supportedCommands["exit"] = cliCommand{
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback:	 commandExit,
		}

		supportedCommands["help"] = cliCommand{
			name: 		 "help",
			description: "Displays a help message",
			callback:	 commandHelp,
		}

		if callbackFunction := supportedCommands[commandName].callback; callbackFunction != nil {
			callbackFunction()
		} else {
			fmt.Println("Unknown command.")
		}

	}
}

func cleanInput(text string) []string {
	loweredInputs := strings.ToLower(text)
	splitInputs := strings.Fields(loweredInputs)
	return splitInputs
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range(supportedCommands) {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
