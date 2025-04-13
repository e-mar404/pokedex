package pokeapi

type Pokedex struct {
	pokemon map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemon: map[string]Pokemon{},
	}
}

func (p *Pokedex) Add(name string, pokemon Pokemon) {
	p.pokemon[name] = pokemon
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	pokemon, ok := p.pokemon[name]
	if !ok {
		return Pokemon{}, false
	}

	return pokemon, true
}
