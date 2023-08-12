package pokeapi

import (
    "fmt"
)

func Pokedex() {
    fmt.Println("Pokedex:")
    for pk := range caughtPokemon {
        fmt.Println("- ", pk)
    }
}
