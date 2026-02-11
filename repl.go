package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "display help message",
		callback:    help,
	},
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() []string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// Scan for the next line of input
		if scanner.Scan() {
			cleanedInput := cleanInput(scanner.Text())
			if len(cleanedInput) == 0 {
				continue
			}
			command := cliCommands[cleanedInput[0]]
			if command.name == "" {
				fmt.Println("Unknown command")
				continue
			}
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}

		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:
help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}
