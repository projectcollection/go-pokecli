package main

import (
	"bufio"
	"fmt"
	pokeapi "github.com/projectcollection/pokecli/internals/pokeapi"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	cb          func() error
}

func helpCmd() error {
	commandString := `
    pokecli:
    xanderjakeq's first go project. just a basic pokecli

    commands:
    `
	for key := range commands {
		command := commands[key]
		commandString = commandString +
			"-" +
			command.name +
			":" +
			command.description +
			"\n"
	}

	fmt.Println(strings.Replace(commandString, "    ", "", -1))
	return nil
}

var commands map[string]command = map[string]command{
	"map": {
		name:        "map",
		description: "list next 20 maps",
		cb: func() error {
			locations, err := pokeapi.Map()

			if err != nil {
				fmt.Println(err)
				return err
			}

			for _, location := range locations {
				fmt.Println(location)
			}

            fmt.Println("")

			return nil
		},
	},
	"mapb": {
		name:        "mapb",
		description: "list previous 20 maps",
		cb: func() error {
			locations, err := pokeapi.Mapb()

			if err != nil {
				fmt.Println(err)
				return err
			}

			for _, location := range locations {
				fmt.Println(location)
			}

            fmt.Println("")

			return nil
		},
	},
	"exit": {
		name:        "exit",
		description: "exit pokecli repl",
		cb: func() error {
			os.Exit(3)
			return nil
		},
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print("pokedex ---> ")
		scanner.Scan()

		text := scanner.Text()

		_, ok := commands[text]

		if text == "help" {
			helpCmd()
			continue
		}

		if !ok {
			fmt.Println("command not found")
			continue
		}

		commands[text].cb()
	}

}
