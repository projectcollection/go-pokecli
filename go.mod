module github.com/projectcollection/pokecli

go 1.21.0

replace github.com/projectcollection/pokecli/internals/pokeapi v0.0.0 => ./internals/pokeapi
replace github.com/projectcollection/pokecli/internals/pokecache v0.0.0 => ./internals/pokecache

require (
    github.com/projectcollection/pokecli/internals/pokeapi v0.0.0
    github.com/projectcollection/pokecli/internals/pokecache v0.0.0
)
