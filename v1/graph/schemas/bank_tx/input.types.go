package account

import (
	"github.com/graphql-go/graphql"
)

var BankTxInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BankTxInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"tx_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"payment_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"payment_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"employee_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
