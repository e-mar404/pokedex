package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	PrevURL string
	NextURL string
}

type ResourceList struct {
	Count    int `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results   []APIResource `json:"results"`
}

type APIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

const BaseURL = "https://pokeapi.co/api/v2"

func List() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in Pokemon",
			Callback:    commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the prev 20 location areas in Pokemon",
			Callback: commandMapb,
		},
	}
}
