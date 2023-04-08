package account

import (
	"github.com/graphql-go/graphql"
)

var PaymentCategoryInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PaymentCategoryInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
