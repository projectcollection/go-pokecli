package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	pokecache "github.com/projectcollection/pokecli/internals/pokecache"
	"io"
	"net/http"
	"time"
)

const (
	baseurl  = "https://pokeapi.co/api/v2"
	location = baseurl + "/location?offset=%d&limit=%d"
	site     = baseurl + "/location-area/%s"
	pokemon  = baseurl + "/pokemon/%s"
	limit    = 20
)

var cache = pokecache.NewCache(5 * time.Minute)

var caughtPokemon = make(map[string]PokemonStats)

func decode[T any](body []byte) (T, error) {
	var data T
	err := json.Unmarshal(body, &data)

	return data, err
}

func getData[T any](api string) (T, error) {
	cachedData, ok := cache.Get(api)

	var z T

	if ok {
		data, err := decode[T](cachedData)
		if err != nil {
			return data, err
		}
	}

	res, err := http.Get(api)

	if err != nil {
		return z, errors.New("something went wrong fetching locations")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return z, err
	}

	if res.StatusCode > 299 {
		return z, errors.New(fmt.Sprintf("failed with status code: %d and body: %s", res.StatusCode, body))
	}

	cache.Add(api, body)
	data, err := decode[T](body)

	if err != nil {
		return z, err
	}

	return data, nil
}
