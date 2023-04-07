package account

import (
	"github.com/graphql-go/graphql"
	transaction "github.com/life-entify/account/v1/graph/schemas/transaction"
)

func DeleteTransaction(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete Transaction",
		Type:        transaction.TransactionType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreateTransaction(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Transaction",
		Type:        transaction.TransactionType,
		Args: graphql.FieldConfigArgument{
			"transaction": &graphql.ArgumentConfig{
				Type: transaction.TransactionInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdateTransaction(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Transaction",
		Type:        transaction.TransactionType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"transaction": &graphql.ArgumentConfig{
				Type: transaction.TransactionInputType,
			},
		},
		Resolve: resolver,
	}
}
