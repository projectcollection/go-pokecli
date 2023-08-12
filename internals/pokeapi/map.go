package pokeapi

import (
	"errors"
	"fmt"
)

type MapData struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func newMapBrowser() (func() ([]string, error), func() ([]string, error)) {
	offset := 0

	//0 for backward, 1 for forward
	lastcall := 1

	return func() ([]string, error) {
			locations := []string{}

			api := fmt.Sprintf(location, offset, limit)
			data, err := getData[MapData](api)

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
			data, err := getData[MapData](api)

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
