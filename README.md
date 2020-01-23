# Pokemon Gopher

Inspired on [https://github.com/lucasbento/graphql-pokemon/](https://github.com/lucasbento/graphql-pokemon/) !!!

Thanks [@lucasbento](https://github.com/lucasbento/)!

<p align="center">
    <img src="https://user-images.githubusercontent.com/578310/72949641-ac4ee980-3d67-11ea-8f44-5bfdfc93717c.png"/>
</p>

## How to use

Simply get Pokémon's information through queries in GraphQL, example:

```graphql
query {
  pokemon(name: "pikachu") {
    id
    number
    name
    attacks {
      special {
        name
        type
        damage
      }
    }
    evolutions {
      id
      number
      name
      weight {
        minimum
        maximum
      }
      attacks {
        fast {
          name
          type
          damage
        }
      }
    }
  }
}
```

## Running

### Production

```sh
go build -o pokemon-gopher
./pokemon-gopher
```

### Development

```sh
go run .
```

## Disclaimer

This was built as part of a talk on GraphQL with GoLang at [@GraphQL-SP](https://www.meetup.com/pt-BR/GraphQL-SP) meetup, check us out, we build cool stuff. ;)