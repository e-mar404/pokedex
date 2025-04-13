package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonAtLocation(area string) (LocationArea, error) {
	url := BaseURL + "/location-area/" + area

	data, ok := c.cache.Get(url)
	if ok {
		var locationArea LocationArea
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)
	// fmt.Println(string(data))

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}
