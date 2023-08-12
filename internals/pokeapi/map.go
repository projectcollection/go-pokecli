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

type Data struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func decode(body []byte) (Data, error) {
	data := Data{}
	err := json.Unmarshal(body, &data)

	return data, err
}

func getData(api string, cachedData []byte, ok bool) (Data, error) {
	if ok {
		data, err := decode(cachedData)
		if err != nil {
			return data, err
		}
	}

	res, err := http.Get(api)

	if err != nil {
		return Data{}, errors.New("something went wrong fetching locations")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return Data{}, err
	}

	if res.StatusCode > 299 {
		return Data{}, errors.New(fmt.Sprintf("failed with status code: %d and body: %s", res.StatusCode, body))
	}

	cache.Add(api, body)
	data, err := decode(body)

	if err != nil {
		return Data{}, err
	}

	return data, nil
}

var cache = pokecache.NewCache(5 * time.Minute)

func newMapBrowser() (func() ([]string, error), func() ([]string, error)) {
	offset := 0

	//0 for backward, 1 for forward
	lastcall := 1

	return func() ([]string, error) {
			locations := []string{}

			api := fmt.Sprintf(location, offset, limit)
			cachedData, ok := cache.Get(api)
			data, err := getData(api, cachedData, ok)

			if err != nil {
				return locations, err
			}

			for _, location := range data.Results {
				locations = append(locations, location.Name)
			}

			offset += limit
			lastcall = 1

			return locations, nil
		},

		func() ([]string, error) {
			locations := []string{}

			if lastcall > 0 {
				offset -= 2 * limit
			} else {
				offset -= limit
			}

			if offset < 0 {
				offset = 0
				return locations, errors.New("no previous locations")
			}

			api := fmt.Sprintf(location, offset, limit)
			cachedData, ok := cache.Get(api)
			data, err := getData(api, cachedData, ok)

			if err != nil {
				return locations, err
			}

			for _, location := range data.Results {
				locations = append(locations, location.Name)
			}

			lastcall = 0

			return locations, nil
		}
}

var forwardMap, backwardMap = newMapBrowser()

func Map() ([]string, error) {
	return forwardMap()
}

func Mapb() ([]string, error) {
	return backwardMap()
}
