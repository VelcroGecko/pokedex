package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	commands map[string]cliCommand
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Poxedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
