package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}

func cleanInput(text string) []string {
	var words []string
	trimString := strings.TrimSpace(text)
	cleanedString := strings.ToLower(trimString)
	words = strings.Split(cleanedString, " ")
	return words
}
