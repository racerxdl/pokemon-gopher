package gqlmodels

import (
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/pokemon-gopher/models"
	"github.com/racerxdl/pokemon-gopher/utils"
)

var PokemonEvolutionRequirement = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PokemonEvolutionRequirement",
	Description: "Represents a Pokémon's requirement to evolve",
	Fields: graphql.Fields{
		"amount": {
			Type:        graphql.Int,
			Description: "The amount of candies to evolve",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.EvolutionRequirements)
				if !ok {
					return nil, utils.TypeError(models.EvolutionRequirements{}, p.Source)
				}
				return d.Amount, nil
			},
		},
		"name": {
			Type:        graphql.String,
			Description: "The name of the candy to evolve",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.EvolutionRequirements)
				if !ok {
					return nil, utils.TypeError(models.EvolutionRequirements{}, p.Source)
				}
				return d.Name, nil
			},
		},
	},
})
