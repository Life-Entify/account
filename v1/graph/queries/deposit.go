package account

import (
	"github.com/graphql-go/graphql"
	payment "github.com/life-entify/account/v1/graph/schemas/payment"
)

func GetDepositByPersonGroup(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Payment By Person ID",
		Type:        payment.DepositAmountByPersons,
		Args: graphql.FieldConfigArgument{
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
func GetPersonDepositBalance(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get user deposit balance",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "PersonDepositBalance",
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
			"person_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
