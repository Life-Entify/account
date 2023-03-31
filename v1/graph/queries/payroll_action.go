package account

import (
	"github.com/graphql-go/graphql"
	payroll_action "github.com/life-entify/account/v1/graph/schemas/payroll_action"
)

func GetPayrollActions(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get PayrollActions with search keyword",
		Type:        graphql.NewList(payroll_action.PayrollActionType),
		Args: graphql.FieldConfigArgument{
			"keyword": &graphql.ArgumentConfig{
				Type: payroll_action.PayrollActionInputType,
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
func GetPayrollActionsByID(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Get PayrollActions with search keyword",
		Type:        payroll_action.PayrollActionType,
		Args: graphql.FieldConfigArgument{
			"_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: resolver,
	}
}
