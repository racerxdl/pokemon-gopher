package gqlmodels

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/pokemon-gopher/database"
	"github.com/racerxdl/pokemon-gopher/models"
	"github.com/racerxdl/pokemon-gopher/utils"
	"regexp"
	"strings"
)

var pkmnNameRegex = regexp.MustCompile(`/[&/\\#,+()$~%.'":*?<>{}]/g`)

var Pokemon = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Pokemon",
	Description: "Represents a Pokémon",
	Fields: graphql.Fields{
		"id": utils.GlobalIdField("Pokemon"),
		"number": {
			Type:        graphql.String,
			Description: "The identifier of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Id[len(d.Id)-3:], nil
			},
		},
		"name": {
			Type:        graphql.String,
			Description: "The name of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Name, nil
			},
		},
		"weight": {
			Type:        PokemonDimension,
			Description: "The minimum and maximum weight of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Weight, nil
			},
		},
		"height": {
			Type:        PokemonDimension,
			Description: "The minimum and maximum weight of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Height, nil
			},
		},
		"classification": {
			Type:        graphql.String,
			Description: "The classification of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Classification, nil
			},
		},
		"types": {
			Type:        graphql.NewList(graphql.String),
			Description: "The type(s) of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Types, nil
			},
		},
		"resistant": {
			Type:        graphql.NewList(graphql.String),
			Description: "The type(s) of Pokémons that this Pokémon is resistant to",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Resistant, nil
			},
		},
		"attacks": {
			Type:        PokemonAttack,
			Description: "The attacks of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Attacks, nil
			},
		},
		"weaknesses": {
			Type:        graphql.NewList(graphql.String),
			Description: "The type(s) of Pokémons that this Pokémon weak to",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.Weaknesses, nil
			},
		},
		"fleeRate": {
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.FleeRate, nil
			},
		},
		"maxCP": {
			Type:        graphql.Int,
			Description: "The maximum CP of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.MaxCP, nil
			},
		},
		"evolutionRequirements": {
			Type:        PokemonEvolutionRequirement,
			Description: "The evolution requirements of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.EvolutionRequirements, nil
			},
		},
		"maxHP": {
			Type:        graphql.Int,
			Description: "The maximum HP of this Pokémon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}
				return d.MaxHP, nil
			},
		},
		"image": {
			Type:        graphql.String,
			Description: "A image of the Pokemon",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				d, ok := p.Source.(models.Pokemon)
				if !ok {
					return nil, utils.TypeError(models.Pokemon{}, p.Source)
				}

				imageName := strings.ToLower(d.Name)
				imageName = pkmnNameRegex.ReplaceAllString(imageName, "")
				imageName = strings.ReplaceAll(imageName, " ", "-")

				return fmt.Sprintf("https://img.pokemondb.net/artwork/%s.jpg", imageName), nil
			},
		},
	},
})

func init() {
	// Added Later due Cyclic Reference
	Pokemon.AddFieldConfig("evolutions", &graphql.Field{
		Type:        graphql.NewList(Pokemon),
		Description: "The evolutions of this Pokémon",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			d, ok := p.Source.(models.Pokemon)
			if !ok {
				return nil, utils.TypeError(models.Pokemon{}, p.Source)
			}
			return database.FetchPokemons(d.Evolutions)
		},
	})
}
