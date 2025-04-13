package pokeapi

type LocationArea struct {
	Id                   int                `json:"id"`
	Name                 string             `json:"string"`
	GameIndex            int                `json:"game_index"`
	EncourterMethodRates []NamedAPIResource `json:"encounter_method_rates"`
	Location             NamedAPIResource   `json:"location"`
	Names                []Name             `json:"names"`
	PokemonEncounters    []struct {
		Pokemon        NamedAPIResource   `json:"pokemon"`
		VersionDetails []NamedAPIResource `json:"version_details"`
	} `json:"pokemon_encounters"`
}
