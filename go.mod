module github.com/projectcollection/pokecli

go 1.21.0

replace github.com/projectcollection/pokecli/internals/pokeapi v0.0.0 => ./internals/

require (
    github.com/projectcollection/pokecli/internals/pokeapi v0.0.0
)
