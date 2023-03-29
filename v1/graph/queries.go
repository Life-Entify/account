package account

import (
	"github.com/graphql-go/graphql"
)

func GetPersons(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Persons with search keyword",
		Type:        graphql.NewList(graphql.String),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
func GetPersonsByID(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Persons with search keyword",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: resolver,
	}
}
