package account

import (
	"github.com/graphql-go/graphql"
	payment_category "github.com/life-entify/account/v1/graph/schemas/payment_category"
)

func GetPaymentCategories(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get GetPaymentCategories with search keyword",
		Type:        graphql.NewList(payment_category.PaymentCategoryType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: payment_category.PaymentCategoryInputType,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
func GetGetPaymentCategoriesByID(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get GetPaymentCategories with search keyword",
		Type:        payment_category.PaymentCategoryType,
		Args: graphql.FieldConfigArgument{
			"_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: resolver,
	}
}
