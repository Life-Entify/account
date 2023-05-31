package account

import (
	"github.com/graphql-go/graphql"
	payment "github.com/life-entify/account/v1/graph/schemas/payment"
)

func GetDepositByPersonGroup(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Payment By Person ID",
		Type:        payment.DepositAmountByPersons,
		Args: graphql.FieldConfigArgument{
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver,
	}
}
func GetPaymentsByEmpAndActionType(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Payment By Employees and Action type",
		Type:        graphql.NewList(payment.PaySummaryByEmployeeType),
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: payment.PaymentInputType,
			},
			"date_filter": &graphql.ArgumentConfig{
				Type: payment.DateStampInputType,
			},
		},
		Resolve: resolver,
	}
}
func GetPayments(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Payments with search keyword",
		Type:        graphql.NewList(payment.PaymentType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: payment.PaymentInputType,
			},
			"date_filter": &graphql.ArgumentConfig{
				Type: payment.DateStampInputType,
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
func GetPaymentsByID(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Payments with search keyword",
		Type:        payment.PaymentType,
		Args: graphql.FieldConfigArgument{
			"_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: resolver,
	}
}
