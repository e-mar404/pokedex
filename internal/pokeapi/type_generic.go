package pokeapi

type ResourceList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}
