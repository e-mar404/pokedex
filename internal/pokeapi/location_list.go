package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(url string) (ResourceList, error) {
	if url == "" {
		url = BaseURL + "/location-area?limit=20&offset=0"
	}

	data, ok := c.cache.Get(url)
	if ok {
		var locationList ResourceList
		if err := json.Unmarshal(data, &locationList); err != nil {
			return ResourceList{}, err
		}

		return locationList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResourceList{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResourceList{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return ResourceList{}, err
	}
	c.cache.Add(url, data)

	var locationList ResourceList
	if err = json.Unmarshal(data, &locationList); err != nil {
		return ResourceList{}, err
	}

	return locationList, nil
}
