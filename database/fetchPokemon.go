package database

import (
	"fmt"
	"github.com/racerxdl/pokemon-gopher/models"
	"strconv"
	"strings"
)

var myFancyDatabase = map[int]models.Pokemon{}
var pkmnList = []models.Pokemon{}
var nameToId = map[string]int{}

func LoadPokemons(pokemons []models.Pokemon) {
	for i, v := range pokemons {
		id, err := v.IntID()
		if err != nil {
			fmt.Printf("Error parsing pokemon entry %d: %s", i, err)
			continue
		}
		myFancyDatabase[id] = v
		pkmnList = append(pkmnList, v)
		nameToId[strings.ToLower(v.Name)] = id
	}
}

func FetchPokemon(pokemon models.IntId) (models.Pokemon, error) {
	pkmn := models.Pokemon{}

	id, err := pokemon.IntID()
	if err != nil {
		return pkmn, err
	}

	pkmn, ok := myFancyDatabase[id]
	if !ok {
		return pkmn, fmt.Errorf("pokemon with id %d not found", id)
	}

	return pkmn, nil
}

func FetchPokemons(pokemons []models.PokemonReference) ([]models.Pokemon, error) {
	var err error
	pkms := make([]models.Pokemon, len(pokemons))

	for i, v := range pokemons {
		pkms[i], err = FetchPokemon(v)
		if err != nil {
			return nil, err
		}
	}

	return pkms, nil
}

func GetPokemons(args map[string]interface{}) ([]models.Pokemon, error) {
	first := args["first"].(int)

	if first <= 0 {
		return nil, fmt.Errorf("first should be > 0")
	}

	return pkmnList[:first], nil
}

func GetPokemonByName(name string) (models.Pokemon, error) {
	pkmn := models.Pokemon{}
	id, found := nameToId[strings.ToLower(name)]
	if !found {
		return pkmn, fmt.Errorf("pokemon with name %s not found", name)
	}

	return myFancyDatabase[id], nil
}

func GetPokemonById(stringId string) (models.Pokemon, error) {
	pkmn := models.Pokemon{}
	id, err := strconv.ParseInt(stringId, 10, 32)
	if err != nil {
		return pkmn, fmt.Errorf("invalid pokemon id: %s", err)
	}

	pkmn, found := myFancyDatabase[int(id)]
	if !found {
		return pkmn, fmt.Errorf("pokemon with id %s not found", stringId)
	}

	return pkmn, nil
}
