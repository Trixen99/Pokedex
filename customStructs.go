package main

type Named struct {
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type LocationArea struct {
	Id                     int                   `json:"id"`
	Name                   string                `json:"name"`
	Game_index             int                   `json:"game_index"`
	Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
	Location               NamedAPIResource      `json:"location"`
	Names                  []Name                `json:"names"`
	Pokemon_encounters     []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	Encounter_method NamedAPIResource          `json:"Encounter_method"`
	Version_details  []EncounterVersionDetails `json:"Version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"Rate"`
	Version NamedAPIResource `json:"Version"`
}

type Name struct {
	Name     string           `json:"Name"`
	Language NamedAPIResource `json:"Language"`
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource         `json:"Pokemon"`
	Version_details []VersionEncounterDetail `json:"Version_details"`
}

type VersionEncounterDetail struct {
	Version           NamedAPIResource `json:"Version"`
	Max_chance        int              `json:"Max_chance"`
	Encounter_details []Encounter      `json:"Encounter_details"`
}

type Encounter struct {
	Min_level        int                `json:"Min_level"`
	Max_level        int                `json:"Max_level"`
	Condition_values []NamedAPIResource `json:"Condition_values"`
	Chance           int                `json:"Chance"`
	Method           NamedAPIResource   `json:"Method"`
}

type Pokemon struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	Base_experience int           `json:"base_experience"`
	Weight          int           `json:"Weight"`
	Height          int           `json:"Height"`
	Stats           []PokemonStat `json:"Stats"`
	Types           []PokemonType `json:"Types"`
}

type PokemonStat struct {
	Stat      NamedAPIResource `json:"stat"`
	Effort    int              `json:"effort"`
	Base_stat int              `json:"base_stat"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"Type"`
}
