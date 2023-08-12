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
	cb          func(args []string) error
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
		cb: func(args []string) error {
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
		cb: func(arg []string) error {
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
	"explore": {
		name:        "explore",
		description: "list previous 20 maps",
		cb: func(arg []string) error {
            site := arg[0]
			encounters, err := pokeapi.Explore(site)

			if err != nil {
				fmt.Println(err)
				return err
			}


            fmt.Println("exploring...")
            fmt.Println("found pokemons:\n")

			for _, encounter := range encounters {
				fmt.Println(encounter.Pokemon.Name)
			}

			fmt.Println("")

			return nil
		},
	},
	"exit": {
		name:        "exit",
		description: "exit pokecli repl",
		cb: func(arg []string) error {
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
		textArr := strings.Split(text, " ")

		token := textArr[0]
		args := textArr[1:]

		command, ok := commands[token]

		if text == "help" {
			helpCmd()
			continue
		}

		if !ok {
			fmt.Println("command not found")
			continue
		}

		command.cb(args)
	}

}
