package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%v: %v", command.name, command.discription)
		fmt.Println()
	}

	return nil
}
