package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scanner.Text()
		splitInput := cleanInput(scanner.Text())

		if len(splitInput) == 0 || splitInput[0] == "" {
			continue
		}

		if cmd, ok := commands[splitInput[0]]; ok {
			cmd.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
