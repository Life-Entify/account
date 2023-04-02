package account

import (
	"github.com/graphql-go/graphql"
)

var PayrollActionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PayrollActionType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.String,
		},
		"is_general": &graphql.Field{
			Type: graphql.String,
		},
		"employee_ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"action_type": &graphql.Field{
			Type: graphql.String,
		},
		"action_kind": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"repeats": &graphql.Field{
			Type: graphql.Int,
		},
		"is_constant": &graphql.Field{
			Type: graphql.String,
		},
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"total_value": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
