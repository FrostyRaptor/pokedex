package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startProgram() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := getCommands()[words[0]]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	list := strings.Fields(lower)
	return list
}

type cliCommand struct {
	name        string
	discription string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			discription: "Shows help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			discription: "Exits the program",
			callback:    commandExit,
		},
	}
}
