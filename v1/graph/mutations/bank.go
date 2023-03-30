package account

import (
	"github.com/graphql-go/graphql"
	bank "github.com/life-entify/account/v1/graph/schemas/bank"
)

func DeleteBank(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete Bank",
		Type:        bank.BankType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreateBank(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Bank",
		Type:        bank.BankType,
		Args: graphql.FieldConfigArgument{
			"bank": &graphql.ArgumentConfig{
				Type: bank.BankInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdateBank(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Bank",
		Type:        bank.BankType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"bank": &graphql.ArgumentConfig{
				Type: bank.BankInputType,
			},
		},
		Resolve: resolver,
	}
}
