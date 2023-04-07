package account

import (
	"context"

	db "github.com/life-entify/account/repository/db"
	"github.com/life-entify/account/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UpdatePayrollAction(ctx context.Context, _id primitive.ObjectID, payrollAction *account.PayrollAction) (*account.PayrollAction, error)
	DeletePayrollAction(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetPayrollActions(ctx context.Context, filterObj *account.PayrollAction, page *db.Pagination) ([]*account.PayrollAction, error)
	CreatePayrollAction(ctx context.Context, payrollAction *account.PayrollAction) (*account.PayrollAction, error)
	ConnectPayrollAction() (*mongo.Client, *mongo.Collection)

	UpdatePayment(ctx context.Context, _id primitive.ObjectID, payment *account.Payment) (*account.Payment, error)
	DeletePayment(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetPayments(ctx context.Context, filterObj *account.Payment, page *db.Pagination) ([]*account.Payment, error)
	CreatePayment(ctx context.Context, payment *account.Payment) (*account.Payment, error)
	ConnectPayment() (*mongo.Client, *mongo.Collection)

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
