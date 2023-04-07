package account

import (
	"github.com/graphql-go/graphql"
)

var PaymentInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PaymentInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pay_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"tx_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"action": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"employee_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"tx_ids": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"total_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"unresolved": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
