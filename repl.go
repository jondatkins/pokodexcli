package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// var cliCommands = map[string]cliCommand{
// 	"exit": {
// 		name:        "exit",
// 		description: "Exit the Pokedex",
// 		callback:    commandExit,
// 	},
// 	"help": {
// 		name:        "help",
// 		description: "display help message",
// 		callback:    help,
// 	},
// }

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
			commandName := cleanedInput[0]
			command, exists := getCommands()[commandName]
			if exists {
				cfg := &config{}
				err := command.callback(cfg)
				fmt.Println("")
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:
help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

const baseURL = "https://pokeapi.co/api/v2/location-area/"

var (
	offset = 0
	limit  = 20
)

func buildURL(offset, limit int) string {
	return fmt.Sprintf("%s?offset=%d&limit=%d", baseURL, offset, limit)
}

func commandMap(cfg *config) error {
	url := buildURL(offset, limit)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	offset += 20
	defer resp.Body.Close()
	var data locationResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if offset-40 < 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	offset -= 40
	url := buildURL(offset, limit)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var data locationResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "go forward through map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "go back through map locations",
			callback:    commandMapb,
		},
	}
}
