package account

import (
	"github.com/graphql-go/graphql"
	emp_db "github.com/life-entify/employee/v1/graph/schemas/employee"
	person "github.com/life-entify/person/v1/graph/schemas"
)

var DepositAmountByPersons = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositAmountByPersons",
	Fields: graphql.Fields{
		"deposit_info": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name: "PersonDepositAmount",
				Fields: graphql.Fields{
					"total_amount": &graphql.Field{
						Type: graphql.Int,
					},
					"_id": &graphql.Field{
						Type: graphql.NewObject(graphql.ObjectConfig{
							Name: "XID",
							Fields: graphql.Fields{
								"action_type": &graphql.Field{
									Type: graphql.String,
								},
								"person_id": &graphql.Field{
									Type: graphql.Int,
								},
							},
						}),
					},
				},
			})),
		},
		"persons": &graphql.Field{
			Type: graphql.NewList(person.PersonType),
		},
	},
})
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
		"employee": &graphql.Field{
			Type: emp_db.EmployeeType,
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
