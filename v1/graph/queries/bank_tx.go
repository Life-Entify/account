package account

import (
	"github.com/graphql-go/graphql"
	bank_tx "github.com/life-entify/account/v1/graph/schemas/bank_tx"
)

func GetBankTxs(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get BankTxs with search keyword",
		Type:        graphql.NewList(bank_tx.BankTxType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: bank_tx.BankTxInputType,
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
func GetBankTxsByID(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get BankTxs with search keyword",
		Type:        bank_tx.BankTxType,
		Args: graphql.FieldConfigArgument{
			"_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: resolver,
	}
}
