package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/racerxdl/pokemon-gopher/database"
	"github.com/racerxdl/pokemon-gopher/models"
	"io/ioutil"
	"net/http"
)

func main() {
	data, err := ioutil.ReadFile("pokemons/pokemons.json")

	if err != nil {
		panic(fmt.Errorf("cannot load pokemons: %s", err))
	}

	var pokemons []models.Pokemon

	err = json.Unmarshal(data, &pokemons)

	if err != nil {
		panic(fmt.Errorf("error parsing pokemon data: %s", err))
	}

	database.LoadPokemons(pokemons)

	schema, err := GetSchema()

	if err != nil {
		panic(fmt.Errorf("error compiling schema: %s", err))
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	fmt.Println("Server running at http://localhost:8080/graphql")

	http.ListenAndServe(":8080", nil)
}
