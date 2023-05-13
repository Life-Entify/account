package account

import (
	"github.com/graphql-go/graphql"
	payment "github.com/life-entify/account/v1/graph/schemas/payment"
	emp_schemas "github.com/life-entify/employee/v1/graph/schemas/employee"
)

func GetPaymentsByEmpAndActionType(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Payment By Employees and Action type",
		Type:        graphql.NewList(payment.PaySummaryByEmployeeType),
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: payment.PaymentInputType,
			},
		},
		Resolve: resolver,
	}
}
func GetPaymentEmployeeIds(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get Distinct Employee Ids for payments",
		Type:        graphql.NewList(emp_schemas.EmployeeType),
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: payment.PaymentInputType,
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
