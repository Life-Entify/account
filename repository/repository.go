package account

import (
	"context"

	db "github.com/life-entify/account/repository/db"
	"github.com/life-entify/account/v1"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UpdateBank(ctx context.Context, _id primitive.ObjectID, bank *account.Bank) (*account.Bank, error)
	DeleteBank(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetBanks(ctx context.Context, filterObj *account.Bank, page *db.Pagination) ([]*account.Bank, error)
	CreateBank(ctx context.Context, bank *account.Bank) (*account.Bank, error)
	ConnectBank() (*mongo.Client, *mongo.Collection)
}
