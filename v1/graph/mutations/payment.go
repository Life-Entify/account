package account

import (
	"github.com/graphql-go/graphql"
	payment "github.com/life-entify/account/v1/graph/schemas/payment"
	transaction "github.com/life-entify/account/v1/graph/schemas/transaction"
)

func DeletePayment(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete Payment",
		Type:        payment.PaymentType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreatePayment(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create Payment",
		Type:        payment.PaymentType,
		Args: graphql.FieldConfigArgument{
			"payment": &graphql.ArgumentConfig{
				Type: payment.PaymentInputType,
			},
			"transactions": &graphql.ArgumentConfig{
				Type: graphql.NewList(transaction.TransactionInputType),
			},
		},
		Resolve: resolver,
	}
}
func UpdatePayment(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update Payment",
		Type:        payment.PaymentType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"payment": &graphql.ArgumentConfig{
				Type: payment.PaymentInputType,
			},
			"transactions": &graphql.ArgumentConfig{
				Type: graphql.NewList(transaction.TransactionInputType),
			},
		},
		Resolve: resolver,
	}
}
