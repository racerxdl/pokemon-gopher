package gqlmodels

import (
	"github.com/graphql-go/graphql"
)

var PokemonEvolutionRequirement = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PokemonEvolutionRequirement",
	Description: "Represents a Pokémon's requirement to evolve",
	Fields: graphql.Fields{
		"amount": {
			Type:        graphql.Int,
			Description: "The amount of candies to evolve",
		},
		"name": {
			Type:        graphql.String,
			Description: "The name of the candy to evolve",
		},
	},
})
