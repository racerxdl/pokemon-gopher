package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

// https://github.com/graphql/graphql-relay-js/blob/35f1926e5a7e345c5aa9c835b9b148d266abd5b9/src/node/node.js#L96
func toGlobalId(typeName, id string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", typeName, id)))
}

// https://github.com/graphql/graphql-relay-js/blob/35f1926e5a7e345c5aa9c835b9b148d266abd5b9/src/node/node.js#L119
func GlobalIdField(name string) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewNonNull(graphql.ID),
		Description: "The ID of the object",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			o := reflect.ValueOf(p.Source)

			f := o.FieldByName("id")
			if !f.IsValid() {
				return toGlobalId(name, f.String()), nil
			}

			f = o.FieldByName("Id")
			if !f.IsValid() {
				return toGlobalId(name, f.String()), nil
			}

			f = o.FieldByName("ID")
			if !f.IsValid() {
				return toGlobalId(name, f.String()), nil
			}

			return toGlobalId(name, ""), nil
		},
	}
}
