package main

import (
	"fmt"
	"bufio"
	"github.com/marcojonkhe1m/pokedexcli/repl"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; ;i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := repl.CleanInput(userInput)
		firstWord := cleanedInput[0]
		fmt.Print("Your command was: ",firstWord, "\n")
	}
}
