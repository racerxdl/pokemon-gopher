package gqlmodels

import (
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/pokemon-gopher/models"
	"github.com/racerxdl/pokemon-gopher/utils"
)

var PokemonDimension = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PokemonDimension",
	Description: "Represents a Pokémon's dimensions",
	Fields: graphql.Fields{
		"minimum": {
			Type:        graphql.String,
			Description: "The minimum value of this dimension",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Dimension)
				if !ok {
					return nil, utils.TypeError(models.Dimension{}, p.Source)
				}
				return d.Minimum, nil
			},
		},
		"maximum": {
			Type:        graphql.String,
			Description: "The maximum value of this dimension",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Dimension)
				if !ok {
					return nil, utils.TypeError(models.Dimension{}, p.Source)
				}
				return d.Maximum, nil
			},
		},
	},
})
