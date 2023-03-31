package account

import (
	"github.com/graphql-go/graphql"
	account "github.com/life-entify/account/v1/graph/schemas/bank"
)

var BankTxType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankTxType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"tx_type": &graphql.Field{
			Type: graphql.String,
		},
		"bank_id": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"payment_type": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"employee_id": &graphql.Field{
			Type: graphql.String,
		},
		"bank": &graphql.Field{
			Type: account.BankType,
		},
	},
})
