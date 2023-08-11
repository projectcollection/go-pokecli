package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

const (
	location = "https://pokeapi.co/api/v2/location?offset=%d&limit=%d"
	limit    = 20
)

func newMapBrowser() (func() ([]string, error), func() ([]string, error)) {
	offset := 0

	//0 for backward, 1 for forward
	lastcall := 1

	return func() ([]string, error) {
			locations := []string{}

			fmt.Println(offset)

			res, err := http.Get(fmt.Sprintf(location, offset, limit))

			if err != nil {
				return locations, errors.New("something went wrong fetching locations")
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()

			if err != nil {
				return locations, err
			}

			if res.StatusCode > 299 {
				return locations, errors.New(fmt.Sprintf("failed with status code: %d and body: %s", res.StatusCode, body))
			}

			data := Data{}
			err = json.Unmarshal(body, &data)

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

			fmt.Println("backing up", offset)

			if offset < 0 {
				offset = 0
				return locations, errors.New("no previous locations")
			}

			res, err := http.Get(fmt.Sprintf(location, offset, limit))

			if err != nil {
				return locations, errors.New("something went wrong fetching locations")
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()

			if err != nil {
				return locations, err
			}

			if res.StatusCode > 299 {
				return locations, errors.New(fmt.Sprintf("failed with status code: %d and body: %s", res.StatusCode, body))
			}

			data := Data{}
			err = json.Unmarshal(body, &data)

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
