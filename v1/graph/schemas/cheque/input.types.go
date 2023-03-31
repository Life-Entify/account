package account

import (
	"github.com/graphql-go/graphql"
)

var ChequeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ChequeInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"cheque_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"cheque_leaflets": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"bank_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"used_leaflets": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
