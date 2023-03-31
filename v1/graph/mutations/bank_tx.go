package account

import (
	"github.com/graphql-go/graphql"
	bank_tx "github.com/life-entify/account/v1/graph/schemas/bank_tx"
)

func DeleteBankTx(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete BankTx",
		Type:        bank_tx.BankTxType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreateBankTx(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create BankTx",
		Type:        bank_tx.BankTxType,
		Args: graphql.FieldConfigArgument{
			"bank_tx": &graphql.ArgumentConfig{
				Type: bank_tx.BankTxInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdateBankTx(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update BankTx",
		Type:        bank_tx.BankTxType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"bank_tx": &graphql.ArgumentConfig{
				Type: bank_tx.BankTxInputType,
			},
		},
		Resolve: resolver,
	}
}
