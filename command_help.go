package main

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
