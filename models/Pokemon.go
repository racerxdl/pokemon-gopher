package models

import "strconv"

type Pokemon struct {
	Id                    string                `json:"id"`
	Name                  string                `json:"name"`
	Classification        string                `json:"classification"`
	Types                 []string              `json:"types"`
	Resistant             []string              `json:"resistant"`
	Weaknesses            []string              `json:"weaknesses"`
	Weight                Dimension             `json:"weight"`
	Height                Dimension             `json:"height"`
	FleeRate              float64               `json:"flee_rate"`
	EvolutionRequirements EvolutionRequirements `json:"evolutionRequirements"`
	Evolutions            []PokemonReference    `json:"evolutions"`
	MaxCP                 int                   `json:"max_cp"`
	MaxHP                 int                   `json:"max_hp"`
	Attacks               PokemonAttack         `json:"attacks"`
}

type PokemonReference struct {
	Id int `json:"id"`
}

type IntId interface {
	IntID() (int, error)
}

func (pkmn Pokemon) IntID() (int, error) {
	id, err := strconv.ParseInt(pkmn.Id, 10, 32)

	return int(id), err
}

func (pkmn PokemonReference) IntID() (int, error) {
	return pkmn.Id, nil
}
