package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(c *Config) error {
	if c.PrevURL == "" {
		fmt.Println("there are no prev areas")
		return nil
	}

	res, err := http.Get(c.PrevURL)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
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
