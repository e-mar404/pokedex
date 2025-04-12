package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(c *Config) error {
	if c.NextURL == "" {
		c.NextURL = BaseURL + "/location-area?limit=20&offset=0" 
	}
	
	res, err := http.Get(c.NextURL)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}
	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("data read error: %w", err)
	}

	var resourceList ResourceList
	if err = json.Unmarshal(data, &resourceList); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	for _, area := range resourceList.Results {
		fmt.Println(area.Name)
	}

	c.PrevURL = resourceList.Previous 
	c.NextURL = resourceList.Next 

	return nil
}
