package account

import (
	"github.com/graphql-go/graphql"
)

func CreatePerson(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Person",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"profile": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func UpdatePerson(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "UpdatePerson Person",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"profile": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
