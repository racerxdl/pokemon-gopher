package gqlmodels

import (
	"github.com/graphql-go/graphql"
)

var PokemonDimension = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PokemonDimension",
	Description: "Represents a Pokémon's dimensions",
	Fields: graphql.Fields{
		"minimum": {
			Type:        graphql.String,
			Description: "The minimum value of this dimension",
		},
		"maximum": {
			Type:        graphql.String,
			Description: "The maximum value of this dimension",
		},
	},
})
