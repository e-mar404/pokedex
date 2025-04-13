package pokeapi

type ResourceList struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []APIResource `json:"results"`
}

type APIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
