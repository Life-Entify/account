package account

import (
	"github.com/graphql-go/graphql"
)

var BankInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BankInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"balance": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"branch": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"is_admin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
