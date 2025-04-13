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

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"string"`
	GameIndex            int                   `json:"game_index"`
	EncourterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             APIResource           `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod APIResource               `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int         `json:"rate"`
	Version APIResource `json:"version"`
}

type Name struct {
	Name     string      `json:"name"`
	Language APIResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        APIResource              `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          APIResource `json:"version"`
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
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	Condition EncounterCondition `json:"condition"`
	Names     []Name             `json:"names"`
}

type EncounterCondition struct {
	Id     int                       `json:"id"`
	Name   string                    `json:"name"`
	Names  []Name                    `json:"names"`
	Values []EncounterConditionValue `json:"values"`
}

type EncounterMethod struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names []Name `json:"names"`
}
