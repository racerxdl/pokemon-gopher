package models

type Attack struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Damage int    `json:"damage"`
}

type PokemonAttack struct {
	Fast    []Attack `json:"fast"`
	Special []Attack `json:"special"`
}
