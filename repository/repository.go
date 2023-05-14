package account

import (
	"context"

	db "github.com/life-entify/account/repository/db"
	"github.com/life-entify/account/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetPaymentByEmpIdAndActionType(ctx context.Context, filterObj *account.Payment, dateFilter *db.DateFilter) ([]*db.PaymentActionTypeSummary, error)
	GetPaymentByEmpIdAndPayType(ctx context.Context, filterObj *account.Payment, dateFilter *db.DateFilter) ([]*db.PaymentTypeSummary, error)
	GetPaymentEmployeeIds(ctx context.Context, filterObj *account.Payment) ([]int64, error)

	UpdatePaymentCategory(ctx context.Context, _id primitive.ObjectID, paymentCategory *account.PaymentCategory) (*account.PaymentCategory, error)
	DeletePaymentCategory(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetPaymentCategories(ctx context.Context, filterObj *account.PaymentCategory, page *db.Pagination) ([]*account.PaymentCategory, error)
	CreatePaymentCategory(ctx context.Context, paymentCategory *account.PaymentCategory) (*account.PaymentCategory, error)
	ConnectPaymentCategory() (*mongo.Client, *mongo.Collection)

	UpdateTransaction(ctx context.Context, _id primitive.ObjectID, transaction *account.Transaction) (*account.Transaction, error)
	DeleteTransaction(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetTransactions(ctx context.Context, filterObj *account.Transaction, page *db.Pagination) ([]*account.Transaction, error)
	GetTransactionsByID(ctx context.Context, _ids []primitive.ObjectID) ([]*account.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *account.Transaction) (*account.Transaction, error)
	ConnectTransaction() (*mongo.Client, *mongo.Collection)

	UpdatePayment(ctx context.Context, _id primitive.ObjectID, payment *account.Payment) (*account.Payment, error)
	DeletePayment(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetPaymentByID(ctx context.Context, _id primitive.ObjectID) (*account.Payment, error)
	GetPayments(ctx context.Context, filterObj *account.Payment, page *db.Pagination) ([]*account.Payment, error)
	CreatePayment(ctx context.Context, payment *account.Payment) (*account.Payment, error)
	ConnectPayment() (*mongo.Client, *mongo.Collection)

	UpdatePayrollAction(ctx context.Context, _id primitive.ObjectID, payrollAction *account.PayrollAction) (*account.PayrollAction, error)
	DeletePayrollAction(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetPayrollActions(ctx context.Context, filterObj *account.PayrollAction, page *db.Pagination) ([]*account.PayrollAction, error)
	CreatePayrollAction(ctx context.Context, payrollAction *account.PayrollAction) (*account.PayrollAction, error)
	ConnectPayrollAction() (*mongo.Client, *mongo.Collection)

	UpdateCheque(ctx context.Context, _id primitive.ObjectID, cheque *account.Cheque) (*account.Cheque, error)
	DeleteCheque(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetCheques(ctx context.Context, filterObj *account.Cheque, page *db.Pagination) ([]*account.Cheque, error)
	CreateCheque(ctx context.Context, cheque *account.Cheque) (*account.Cheque, error)
	ConnectCheque() (*mongo.Client, *mongo.Collection)

	UpdateBankTx(ctx context.Context, _id primitive.ObjectID, bank *account.BankTx) (*account.BankTx, error)
	DeleteBankTx(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetBankTxs(ctx context.Context, filterObj *account.BankTx, page *db.Pagination) ([]*account.BankTx, error)
	CreateBankTx(ctx context.Context, bank *account.BankTx) (*account.BankTx, error)
	ConnectBankTx() (*mongo.Client, *mongo.Collection)

	UpdateBank(ctx context.Context, _id primitive.ObjectID, bank *account.Bank) (*account.Bank, error)
	DeleteBank(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetBanks(ctx context.Context, filterObj *account.Bank, page *db.Pagination) ([]*account.Bank, error)
	CreateBank(ctx context.Context, bank *account.Bank) (*account.Bank, error)
	ConnectBank() (*mongo.Client, *mongo.Collection)
}
