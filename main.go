package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var words []string
	trimString := strings.TrimSpace(text)
	cleanedString := strings.ToLower(trimString)
	words = strings.Split(cleanedString, " ")
	return words
}
