package account

import (
	"github.com/graphql-go/graphql"
)

var PayrollActionInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PayrollActionInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"is_general": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"employee_ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
		"action_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"action_kind": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"repeats": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"is_constant": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"count": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"total_value": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
