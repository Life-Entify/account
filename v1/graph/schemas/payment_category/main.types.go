package account

import (
	"github.com/graphql-go/graphql"
)

var PaymentCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentCategoryType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})
