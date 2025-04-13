package pokeapi

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"string"`
	GameIndex            int                   `json:"game_index"`
	EncourterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource           `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource               `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int         `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource              `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int         `json:"max_chance"`
	EncounterDetails []Encounter `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int                     `json:"min_level"`
	MaxLevel        int                     `json:"max_level"`
	ConditionValues []EncounterConditionValue `json:"condition_values"`
	Chance          int                     `json:"chance"`
	Method          EncounterMethod         `json:"method"`
}

type EncounterConditionValue struct {
	NamedAPIResource
	Condition EncounterCondition `json:"condition"`
	Names     []Name             `json:"names"`
}

type EncounterCondition struct {
	NamedAPIResource
	Names  []Name                    `json:"names"`
	Values []EncounterConditionValue `json:"values"`
}

type EncounterMethod struct {
	NamedAPIResource
	Order int    `json:"order"`
	Names []Name `json:"names"`
}
