package account

import (
	"github.com/graphql-go/graphql"
)

var PaySummaryByEmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaySummaryByEmployeeType",
	Fields: graphql.Fields{
		"action_types": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name: "PayActionTypesType",
				Fields: graphql.Fields{
					"action_type": &graphql.Field{
						Type: graphql.String,
					},
					"total_amount": &graphql.Field{
						Type: graphql.Int,
					},
				},
			})),
		},
		"pay_types": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name: "PayTypesType",
				Fields: graphql.Fields{
					"pay_type": &graphql.Field{
						Type: graphql.String,
					},
					"tx_type": &graphql.Field{
						Type: graphql.String,
					},
					"total_amount": &graphql.Field{
						Type: graphql.Int,
					},
				},
			})),
		},
		"employee_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
var PaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentType",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"pay_type": &graphql.Field{
			Type: graphql.String,
		},
		"tx_type": &graphql.Field{
			Type: graphql.String,
		},
		"action_type": &graphql.Field{
			Type: graphql.String,
		},
		"bank_id": &graphql.Field{
			Type: graphql.String,
		},
		"cheque_id": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"client": &graphql.Field{
			Type: graphql.String,
		},
		"person_id": &graphql.Field{
			Type: graphql.Int,
		},
		"employee_id": &graphql.Field{
			Type: graphql.Int,
		},
		"tx_ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"total_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"unresolved": &graphql.Field{
			Type: graphql.String,
		},
	},
})
