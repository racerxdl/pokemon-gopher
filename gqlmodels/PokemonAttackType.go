package gqlmodels

import (
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/pokemon-gopher/models"
	"github.com/racerxdl/pokemon-gopher/utils"
)

var AttackType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Attack",
	Description: "Represents a Pokémon's attack types",
	Fields: graphql.Fields{
		"name": {
			Type:        graphql.String,
			Description: "The name of this Pokémon attack",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Attack)
				if !ok {
					return nil, utils.TypeError(models.Attack{}, p.Source)
				}
				return d.Name, nil
			},
		},
		"type": {
			Type:        graphql.String,
			Description: "The type of this Pokémon attack",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Attack)
				if !ok {
					return nil, utils.TypeError(models.Attack{}, p.Source)
				}
				return d.Type, nil
			},
		},
		"damage": {
			Type:        graphql.Int,
			Description: "The damage of this Pokémon attack",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Attack)
				if !ok {
					return nil, utils.TypeError(models.Attack{}, p.Source)
				}
				return d.Damage, nil
			},
		},
	},
})

var PokemonAttack = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PokemonAttack",
	Description: "Represents a Pokemon's attack types",
	Fields: graphql.Fields{
		"fast": {
			Type:        graphql.NewList(AttackType),
			Description: "The fast attacks of this pokemon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.PokemonAttack)
				if !ok {
					return nil, utils.TypeError(models.PokemonAttack{}, p.Source)
				}
				return d.Fast, nil
			},
		},
		"special": {
			Type:        graphql.NewList(AttackType),
			Description: "The special attacks of this pokemon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.PokemonAttack)
				if !ok {
					return nil, utils.TypeError(models.PokemonAttack{}, p.Source)
				}
				return d.Special, nil
			},
		},
	},
})
