package account

import (
	"github.com/graphql-go/graphql"
	payroll_action "github.com/life-entify/account/v1/graph/schemas/payroll_action"
)

func DeletePayrollAction(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Delete PayrollAction",
		Type:        payroll_action.PayrollActionType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver,
	}
}
func CreatePayrollAction(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Create PayrollAction",
		Type:        payroll_action.PayrollActionType,
		Args: graphql.FieldConfigArgument{
			"payroll_action": &graphql.ArgumentConfig{
				Type: payroll_action.PayrollActionInputType,
			},
		},
		Resolve: resolver,
	}
}
func UpdatePayrollAction(resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Description: "Update PayrollAction",
		Type:        payroll_action.PayrollActionType,
		Args: graphql.FieldConfigArgument{
			"_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"payroll_action": &graphql.ArgumentConfig{
				Type: payroll_action.PayrollActionInputType,
			},
		},
		Resolve: resolver,
	}
}
