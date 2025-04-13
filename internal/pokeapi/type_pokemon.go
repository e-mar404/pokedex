package pokeapi

type Pokemon struct {
	Id             int                `json:"id"`
	Name           string             `json:"name"`
	BaseExperience int                `json:"base_experience"`
	Height         int                `json:"height"`
	IsDefault      bool               `json:"is_default"`
	Order          int                `json:"order"`
	Weight         int                `json:"weight"`
	Abilities      []PokemonAbility   `json:"abilities"`
	Forms          []NamedAPIResource `json:"forms"`
	GamesIndices   []struct {
		GameIndex int              `json:"game_index"`
		Version   NamedAPIResource `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           NamedAPIResource `json:"item"`
		VersionDetails []struct {
			Version NamedAPIResource `json:"version"`
			Rarity  int              `json:"rarity"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                NamedAPIResource `json:"move"`
		VersionGroupDetails []struct {
			MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
			VersionGrouyp   NamedAPIResource `json:"version_group"`
			LevelLearnedAt  int              `json:"level_learned_at"`
			Order           int              `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	PastTypes []struct {
		Generation NamedAPIResource `json:"generation"`
		Types      []PokemonType    `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation NamedAPIResource `json:"generation"`
		Abilities  []PokemonAbility `json:"abilities"`
	} `json:"past_abilities"`
	Sprites struct {
		FrontDefault     string `json:"front_default"`
		FrontShiny       string `json:"front_shiny"`
		FrontFemale      string `json:"front_female"`
		FrontShinyFemale string `json:"front_shiny_female"`
		BackDefault      string `json:"back_default"`
		BackShiny        string `json:"back_shiny"`
		BackFemale       string `json:"back_female"`
		BackShinyFemale  string `json:"back_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Species NamedAPIResource `json:"species"`
	Stats   []struct {
		Stat     NamedAPIResource `json:"stat"`
		Effort   int              `json:"effort"`
		BaseStat int              `json:"base_stat"`
	} `json:"stats"`
	Types []PokemonType `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}
