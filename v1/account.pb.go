// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: v1/account.proto

package account

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TxType int32

const (
	TxType_INCOME      TxType = 0
	TxType_EXPENDITURE TxType = 1
)

// Enum value maps for TxType.
var (
	TxType_name = map[int32]string{
		0: "INCOME",
		1: "EXPENDITURE",
	}
	TxType_value = map[string]int32{
		"INCOME":      0,
		"EXPENDITURE": 1,
	}
)

func (x TxType) Enum() *TxType {
	p := new(TxType)
	*p = x
	return p
}

func (x TxType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_account_proto_enumTypes[0].Descriptor()
}

func (TxType) Type() protoreflect.EnumType {
	return &file_v1_account_proto_enumTypes[0]
}

func (x TxType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxType.Descriptor instead.
func (TxType) EnumDescriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{0}
}

type AccountAction int32

const (
	AccountAction_RECEIVE_PAY        AccountAction = 0
	AccountAction_PAY                AccountAction = 1
	AccountAction_RECEIVE_DEPOSIT    AccountAction = 2
	AccountAction_DEPOSIT_WITHDRAWAL AccountAction = 3
	AccountAction_REGISTER_CREDIT    AccountAction = 4
	AccountAction_REDEEM_CREDIT      AccountAction = 5
)

// Enum value maps for AccountAction.
var (
	AccountAction_name = map[int32]string{
		0: "RECEIVE_PAY",
		1: "PAY",
		2: "RECEIVE_DEPOSIT",
		3: "DEPOSIT_WITHDRAWAL",
		4: "REGISTER_CREDIT",
		5: "REDEEM_CREDIT",
	}
	AccountAction_value = map[string]int32{
		"RECEIVE_PAY":        0,
		"PAY":                1,
		"RECEIVE_DEPOSIT":    2,
		"DEPOSIT_WITHDRAWAL": 3,
		"REGISTER_CREDIT":    4,
		"REDEEM_CREDIT":      5,
	}
)

func (x AccountAction) Enum() *AccountAction {
	p := new(AccountAction)
	*p = x
	return p
}

func (x AccountAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountAction) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_account_proto_enumTypes[1].Descriptor()
}

func (AccountAction) Type() protoreflect.EnumType {
	return &file_v1_account_proto_enumTypes[1]
}

func (x AccountAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountAction.Descriptor instead.
func (AccountAction) EnumDescriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{1}
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId         string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	PayType     string   `protobuf:"bytes,2,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	TxType      string   `protobuf:"bytes,3,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	Action      string   `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Description string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PersonId    int64    `protobuf:"varint,6,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	EmployeeId  int64    `protobuf:"varint,7,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	TxIds       []string `protobuf:"bytes,8,rep,name=tx_ids,json=txIds,proto3" json:"tx_ids,omitempty"`
	TotalAmount int64    `protobuf:"varint,9,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	CreatedAt   string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Unresolved  string   `protobuf:"bytes,11,opt,name=unresolved,proto3" json:"unresolved,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Payment) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

func (x *Payment) GetTxType() string {
	if x != nil {
		return x.TxType
	}
	return ""
}

func (x *Payment) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Payment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Payment) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *Payment) GetEmployeeId() int64 {
	if x != nil {
		return x.EmployeeId
	}
	return 0
}

func (x *Payment) GetTxIds() []string {
	if x != nil {
		return x.TxIds
	}
	return nil
}

func (x *Payment) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Payment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Payment) GetUnresolved() string {
	if x != nil {
		return x.Unresolved
	}
	return ""
}

type PayrollAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId         string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Active      string   `protobuf:"bytes,4,opt,name=active,proto3" json:"active,omitempty"`
	IsGeneral   string   `protobuf:"bytes,5,opt,name=is_general,json=isGeneral,proto3" json:"is_general,omitempty"`
	EmployeeIds []string `protobuf:"bytes,6,rep,name=employee_ids,json=employeeIds,proto3" json:"employee_ids,omitempty"`
	ActionType  string   `protobuf:"bytes,7,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	ActionKind  string   `protobuf:"bytes,8,opt,name=action_kind,json=actionKind,proto3" json:"action_kind,omitempty"`
	Amount      int64    `protobuf:"varint,9,opt,name=amount,proto3" json:"amount,omitempty"`
	Repeats     int64    `protobuf:"varint,10,opt,name=repeats,proto3" json:"repeats,omitempty"`
	IsConstant  string   `protobuf:"bytes,11,opt,name=is_constant,json=isConstant,proto3" json:"is_constant,omitempty"`
	Count       int64    `protobuf:"varint,12,opt,name=count,proto3" json:"count,omitempty"`
	TotalValue  int64    `protobuf:"varint,13,opt,name=total_value,json=totalValue,proto3" json:"total_value,omitempty"`
}

func (x *PayrollAction) Reset() {
	*x = PayrollAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayrollAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayrollAction) ProtoMessage() {}

func (x *PayrollAction) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayrollAction.ProtoReflect.Descriptor instead.
func (*PayrollAction) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *PayrollAction) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *PayrollAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PayrollAction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PayrollAction) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

func (x *PayrollAction) GetIsGeneral() string {
	if x != nil {
		return x.IsGeneral
	}
	return ""
}

func (x *PayrollAction) GetEmployeeIds() []string {
	if x != nil {
		return x.EmployeeIds
	}
	return nil
}

func (x *PayrollAction) GetActionType() string {
	if x != nil {
		return x.ActionType
	}
	return ""
}

func (x *PayrollAction) GetActionKind() string {
	if x != nil {
		return x.ActionKind
	}
	return ""
}

func (x *PayrollAction) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayrollAction) GetRepeats() int64 {
	if x != nil {
		return x.Repeats
	}
	return 0
}

func (x *PayrollAction) GetIsConstant() string {
	if x != nil {
		return x.IsConstant
	}
	return ""
}

func (x *PayrollAction) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PayrollAction) GetTotalValue() int64 {
	if x != nil {
		return x.TotalValue
	}
	return 0
}

type Cheque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId            string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	ChequeNumber   string `protobuf:"bytes,2,opt,name=cheque_number,json=chequeNumber,proto3" json:"cheque_number,omitempty"`
	ChequeLeaflets int64  `protobuf:"varint,3,opt,name=cheque_leaflets,json=chequeLeaflets,proto3" json:"cheque_leaflets,omitempty"`
	BankId         string `protobuf:"bytes,4,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
	Description    string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	UsedLeaflets   int64  `protobuf:"varint,6,opt,name=used_leaflets,json=usedLeaflets,proto3" json:"used_leaflets,omitempty"`
	CreatedAt      string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Bank           *Bank  `protobuf:"bytes,8,opt,name=bank,proto3" json:"bank,omitempty"`
}

func (x *Cheque) Reset() {
	*x = Cheque{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cheque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cheque) ProtoMessage() {}

func (x *Cheque) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cheque.ProtoReflect.Descriptor instead.
func (*Cheque) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *Cheque) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Cheque) GetChequeNumber() string {
	if x != nil {
		return x.ChequeNumber
	}
	return ""
}

func (x *Cheque) GetChequeLeaflets() int64 {
	if x != nil {
		return x.ChequeLeaflets
	}
	return 0
}

func (x *Cheque) GetBankId() string {
	if x != nil {
		return x.BankId
	}
	return ""
}

func (x *Cheque) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cheque) GetUsedLeaflets() int64 {
	if x != nil {
		return x.UsedLeaflets
	}
	return 0
}

func (x *Cheque) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Cheque) GetBank() *Bank {
	if x != nil {
		return x.Bank
	}
	return nil
}

type Bank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId         string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Bank        string `protobuf:"bytes,2,opt,name=bank,proto3" json:"bank,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Number      string `protobuf:"bytes,4,opt,name=number,proto3" json:"number,omitempty"`
	Branch      string `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	IsAdmin     string `protobuf:"bytes,7,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	Active      string `protobuf:"bytes,8,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Bank) Reset() {
	*x = Bank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bank) ProtoMessage() {}

func (x *Bank) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bank.ProtoReflect.Descriptor instead.
func (*Bank) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *Bank) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Bank) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *Bank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bank) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Bank) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *Bank) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Bank) GetIsAdmin() string {
	if x != nil {
		return x.IsAdmin
	}
	return ""
}

func (x *Bank) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

type BankTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId         string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	TxType      string `protobuf:"bytes,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	BankId      string `protobuf:"bytes,3,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
	Amount      int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PaymentType string `protobuf:"bytes,6,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	EmployeeId  string `protobuf:"bytes,8,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	Bank        *Bank  `protobuf:"bytes,9,opt,name=bank,proto3" json:"bank,omitempty"`
}

func (x *BankTx) Reset() {
	*x = BankTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankTx) ProtoMessage() {}

func (x *BankTx) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankTx.ProtoReflect.Descriptor instead.
func (*BankTx) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *BankTx) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *BankTx) GetTxType() string {
	if x != nil {
		return x.TxType
	}
	return ""
}

func (x *BankTx) GetBankId() string {
	if x != nil {
		return x.BankId
	}
	return ""
}

func (x *BankTx) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BankTx) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BankTx) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *BankTx) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BankTx) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *BankTx) GetBank() *Bank {
	if x != nil {
		return x.Bank
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId        string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	TxType     string `protobuf:"bytes,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	Amount     int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentId  int64  `protobuf:"varint,4,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Remark     string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt  string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CategoryId string `protobuf:"bytes,8,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_v1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Transaction) GetTxType() string {
	if x != nil {
		return x.TxType
	}
	return ""
}

func (x *Transaction) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *Transaction) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Transaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Transaction) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

var File_v1_account_proto protoreflect.FileDescriptor

var file_v1_account_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xbf, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x78, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x6e,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x22, 0xfc, 0x02, 0x0a, 0x0d, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x06, 0x43, 0x68, 0x65, 0x71,
	0x75, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x71,
	0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x71,
	0x75, 0x65, 0x5f, 0x6c, 0x65, 0x61, 0x66, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x4c, 0x65, 0x61, 0x66, 0x6c, 0x65, 0x74,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x66, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x66, 0x6c, 0x65, 0x74,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0xc4,
	0x01, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6b, 0x54, 0x78,
	0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0xc6,
	0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x2a, 0x25, 0x0a, 0x06, 0x54, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x45, 0x58, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x54, 0x55, 0x52, 0x45, 0x10, 0x01, 0x2a, 0x7e,
	0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x41, 0x59, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x43,
	0x45, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52,
	0x41, 0x57, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x52,
	0x45, 0x44, 0x45, 0x45, 0x4d, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x05, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_account_proto_rawDescOnce sync.Once
	file_v1_account_proto_rawDescData = file_v1_account_proto_rawDesc
)

func file_v1_account_proto_rawDescGZIP() []byte {
	file_v1_account_proto_rawDescOnce.Do(func() {
		file_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_account_proto_rawDescData)
	})
	return file_v1_account_proto_rawDescData
}

var file_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_account_proto_goTypes = []interface{}{
	(TxType)(0),           // 0: v1.TxType
	(AccountAction)(0),    // 1: v1.AccountAction
	(*Payment)(nil),       // 2: v1.Payment
	(*PayrollAction)(nil), // 3: v1.PayrollAction
	(*Cheque)(nil),        // 4: v1.Cheque
	(*Bank)(nil),          // 5: v1.Bank
	(*BankTx)(nil),        // 6: v1.BankTx
	(*Transaction)(nil),   // 7: v1.Transaction
}
var file_v1_account_proto_depIdxs = []int32{
	5, // 0: v1.Cheque.bank:type_name -> v1.Bank
	5, // 1: v1.BankTx.bank:type_name -> v1.Bank
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_account_proto_init() }
func file_v1_account_proto_init() {
	if File_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayrollAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cheque); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bank); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankTx); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_account_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_account_proto_goTypes,
		DependencyIndexes: file_v1_account_proto_depIdxs,
		EnumInfos:         file_v1_account_proto_enumTypes,
		MessageInfos:      file_v1_account_proto_msgTypes,
	}.Build()
	File_v1_account_proto = out.File
	file_v1_account_proto_rawDesc = nil
	file_v1_account_proto_goTypes = nil
	file_v1_account_proto_depIdxs = nil
}
