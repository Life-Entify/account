package account

import (
	"github.com/graphql-go/graphql"
)

func GetPersonDepositBalance(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get user deposit balance",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "PaymentType",
			Fields: graphql.Fields{
				"deposit": &graphql.Field{
					Type: graphql.Int,
				},
				"used": &graphql.Field{
					Type: graphql.Int,
				},
				"withdrawn": &graphql.Field{
					Type: graphql.Int,
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"personId": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
