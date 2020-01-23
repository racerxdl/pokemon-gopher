package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/pokemon-gopher/database"
	"github.com/racerxdl/pokemon-gopher/gqlmodels"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Query any pokemon by number or name",
	Fields: graphql.Fields{
		"pokemons": {
			Type: graphql.NewList(gqlmodels.Pokemon),
			Args: graphql.FieldConfigArgument{
				"first": {
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return database.GetPokemons(p.Args)
			},
		},
		"pokemon": {
			Type: gqlmodels.Pokemon,
			Args: graphql.FieldConfigArgument{
				"id": {
					Type: graphql.String,
				},
				"name": {
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Args["id"] != nil {
					id := p.Args["id"].(string)
					return database.GetPokemonById(id)
				}

				if p.Args["name"] != nil {
					name := p.Args["name"].(string)
					return database.GetPokemonByName(name)
				}

				return nil, fmt.Errorf("you need to specify \"id\" or \"name\" of the pokemon")
			},
		},
	},
})

func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryType,
	})
}
