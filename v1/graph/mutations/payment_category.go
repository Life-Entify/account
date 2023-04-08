package account

import (
	"github.com/graphql-go/graphql"
	payment_category "github.com/life-entify/account/v1/graph/schemas/payment_category"
)

func DeletePaymentCategory(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete PaymentCategory",
		Type:        payment_category.PaymentCategoryType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreatePaymentCategory(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create PaymentCategory",
		Type:        payment_category.PaymentCategoryType,
		Args: graphql.FieldConfigArgument{
			"payment_category": &graphql.ArgumentConfig{
				Type: payment_category.PaymentCategoryInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdatePaymentCategory(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update PaymentCategory",
		Type:        payment_category.PaymentCategoryType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"payment_category": &graphql.ArgumentConfig{
				Type: payment_category.PaymentCategoryInputType,
			},
		},
		Resolve: resolver,
	}
}
