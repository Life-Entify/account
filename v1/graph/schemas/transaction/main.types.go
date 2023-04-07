package account

import (
	"github.com/graphql-go/graphql"
)

var TransactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TransactionType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"tx_type": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.String,
		},
		"payment_id": &graphql.Field{
			Type: graphql.String,
		},
		"remark": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"category_id": &graphql.Field{
			Type: graphql.String,
		},
	},
})
