package main

import "strings"

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	list := strings.Fields(lower)
	return list
}
