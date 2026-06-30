package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		var userInput string
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}

		userInput = scanner.Text()
		cleanUserInput := cleanInput(userInput)
		fmt.Println("Your command was:", cleanUserInput[0])
	}
}
