package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
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

func Map(dir string) ([]string, error) {
	locations := []string{}

	if dir != "<" && dir != ">" {
		return locations, errors.New("dir only accepts '<' or '>'")
	}

	offset := 0

	if dir == ">" {
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
			return locations, errors.New(fmt.Sprint("failed with status code: %d and body: %s", res.StatusCode, body))
		}

		data := Data{}

		err = json.Unmarshal(body, &data)

		if err != nil {
			return locations, err
		}

        for _, location := range data.Results {
            locations = append(locations, location.Name)
        }

        return locations, nil
	}

    return locations, nil
}
