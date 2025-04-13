package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon string) (Pokemon, bool, error) {
	url := fmt.Sprintf("%s/pokemon/%s", BaseURL, pokemon)

	data, ok := c.cache.Get(url)
	if ok {
		var p Pokemon
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, false, err
		}

		caught := tryCatchPokemon(p)

		return p, caught, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, false, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, false, err
	}
	defer res.Body.Close()

	var p Pokemon
	data, err = io.ReadAll(res.Body)
	if err := json.Unmarshal(data, &p); err != nil {
		return Pokemon{}, false, err
	}

	caught := tryCatchPokemon(p)

	return p, caught, nil
}

func tryCatchPokemon(p Pokemon) bool {
	// the pokemon with the highest base experience are Arceus, Happiny, Chansey and Blissey with 255, so to make those those hard to catch i will make it a percentage of 1 - base exp/ 270
	// check Trivia section: https://bulbapedia.bulbagarden.net/wiki/Experience
	catchProbability := (1.0 - (float64(p.BaseExperience) / 270.0))
	randomNumber := rand.Float64()
	if randomNumber > catchProbability {
		return false
	}
	return true
}
