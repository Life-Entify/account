package account

import (
	"github.com/graphql-go/graphql"
)

var BankType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BankType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"bank": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"number": &graphql.Field{
			Type: graphql.String,
		},
		"branch": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"balance": &graphql.Field{
			Type: graphql.Int,
		},
		"is_admin": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.String,
		},
	},
})
