package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanTextLst := cleanInput(text)
		fmt.Printf("Your command was: %s\n", cleanTextLst[0])
	}
}
