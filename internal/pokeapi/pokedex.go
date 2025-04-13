package pokeapi

type Pokedex map[string]Pokemon

func NewPokedex() Pokedex {
	return Pokedex{}
}
