package account

import (
	"github.com/graphql-go/graphql"
)

var TransactionInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TransactionInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"tx_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"payment_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"remark": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"category_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
