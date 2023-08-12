package pokeapi

import (
	"fmt"
	"reflect"
	"slices"
)

var fieldsToShow = []string{"Height", "Weight", "Stats", "Types", "Abilities"}

func Inspect(pk string) {

	pokemon, ok := caughtPokemon[pk]

	if !ok {
		fmt.Println(fmt.Sprintf("haven't caught a %s yet", pk))
		return
	}

	resString := ""

	fields := reflect.ValueOf(pokemon)
	structType := fields.Type()
	numFields := fields.NumField()

	for i := 0; i < numFields; i++ {
		field := structType.Field(i)

		if !slices.Contains(fieldsToShow, field.Name) {
			continue
		}
		fieldVal := fields.Field(i)

		switch fieldVal.Kind() {
		case reflect.Int:
			fmt.Printf("%s: %v\n", field.Name, fieldVal)

		case reflect.Slice:
			switch field.Name {
			case "Abilities":
				vals := pokemon.Abilities

				fmt.Println("Abilities:")
				for _, val := range vals {
					fmt.Println("- ", val.Ability.Name)
				}

			case "Types":
				vals := pokemon.Types

				fmt.Println("Types:")
				for _, val := range vals {
					fmt.Println("- ", val.Type.Name)
				}
			case "Stats":
				vals := pokemon.Stats

				fmt.Println("Stats:")
				for _, val := range vals {
					fmt.Println("- ", val.Stat.Name)
				}
			}
		}
	}

	fmt.Println(resString)
}
