package pokeapi

type Pokedex struct {
	Pokemon map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokemon: map[string]Pokemon{},
	}
}

func (p *Pokedex) Add(name string, pokemon Pokemon) {
	p.Pokemon[name] = pokemon
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	pokemon, ok := p.Pokemon[name]
	if !ok {
		return Pokemon{}, false
	}

	return pokemon, true
}
