package commands

import "fmt"

func mapf(c *Config) error {
	locationList, err := c.PokeClient.ListLocations(c.nextURL)
	if err != nil {
		return err
	}

	for _, area := range locationList.Results {
		fmt.Println(area.Name)
	}

	c.prevURL = locationList.Previous 
	c.nextURL = locationList.Next 

	return nil
}

func mapb(c *Config) error {
	locationList, err := c.PokeClient.ListLocations(c.prevURL)
	if err != nil {
		return err
	}

	for _, area := range locationList.Results {
		fmt.Println(area.Name)
	}

	c.prevURL = locationList.Previous 
	c.nextURL = locationList.Next 

	return nil
}
