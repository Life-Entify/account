package account

import (
	"github.com/graphql-go/graphql"
	bank "github.com/life-entify/account/v1/graph/schemas/bank"
)

var ChequeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ChequeType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"cheque_number": &graphql.Field{
			Type: graphql.String,
		},
		"cheque_leaflets": &graphql.Field{
			Type: graphql.Int,
		},
		"bank_id": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"used_leaflets": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"bank": &graphql.Field{
			Type: bank.BankType,
		},
	},
})
