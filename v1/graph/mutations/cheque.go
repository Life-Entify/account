package account

import (
	"github.com/graphql-go/graphql"
	cheque "github.com/life-entify/account/v1/graph/schemas/cheque"
)

func DeleteCheque(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete Cheque",
		Type:        cheque.ChequeType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreateCheque(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Cheque",
		Type:        cheque.ChequeType,
		Args: graphql.FieldConfigArgument{
			"cheque": &graphql.ArgumentConfig{
				Type: cheque.ChequeInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdateCheque(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Cheque",
		Type:        cheque.ChequeType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"cheque": &graphql.ArgumentConfig{
				Type: cheque.ChequeInputType,
			},
		},
		Resolve: resolver,
	}
}
