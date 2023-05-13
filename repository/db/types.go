package account

type Pagination struct {
	Limit, Skip int64
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
	EmployeeId  int64                    `json:"employee_id"`
	PayTypes    []*EmpPayTypeTotal       `json:"pay_types"`
	ActionTypes []*EmpPayActionTypeTotal `json:"action_types"`
}
