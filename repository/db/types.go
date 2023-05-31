package account

import "github.com/life-entify/employee/v1"

type Pagination struct {
	Limit, Skip int64
}
type DateFilter struct {
	DateStampFrom string `json:"date_stamp_from"`
	DateStampTo   string `json:"date_stamp_to"`
}
type DepositSummary struct {
	XId struct {
		PersonId   int    `json:"person_id"`
		ActionType string `json:"action_type"`
	} `json:"_id"`
	TotalAmount int `json:"total_amount"`
}
type PaymentActionTypeSummary struct {
	XId struct {
		EmployeeId int    `json:"employee_id"`
		ActionType string `json:"action_type"`
	} `json:"_id"`
	TotalAmount int `json:"total_amount"`
}
type PaymentTypeSummary struct {
	XId struct {
		EmployeeId int    `json:"employee_id"`
		PayType    string `json:"pay_type"`
		TxType     string `json:"tx_type"`
	} `json:"_id"`
	TotalAmount int `json:"total_amount"`
}
type EmpPayTypeTotal struct {
	TxType      string `json:"tx_type"`
	PayType     string `json:"pay_type"`
	TotalAmount int    `json:"total_amount"`
}
type EmpPayActionTypeTotal struct {
	ActionType  string `json:"action_type"`
	TotalAmount int    `json:"total_amount"`
}
type PaySummaryByEmployee struct {
	Employee    *employee.Employee       `json:"employee"`
	PayTypes    []*EmpPayTypeTotal       `json:"pay_types"`
	ActionTypes []*EmpPayActionTypeTotal `json:"action_types"`
}
