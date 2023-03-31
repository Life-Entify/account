package account

import (
	"context"
	"reflect"

	common "github.com/life-entify/account/common"
	errors "github.com/life-entify/account/errors"
	db "github.com/life-entify/account/repository/db"
	"github.com/life-entify/account/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	BANK_TX_COLLECTION = "bankTx_txs"
)

func (db *MongoDB) ConnectBankTx() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(BANK_TX_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdateBankTx(ctx context.Context, _id primitive.ObjectID, bankTx *account.BankTx) (*account.BankTx, error) {
	client, coll := db.ConnectBankTx()
	defer MongoDisconnect(client)

	var jsonBankTx bson.M
	common.ToJSONStruct(bankTx, &jsonBankTx)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonBankTx})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result    bson.M
		newBankTx account.BankTx
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newBankTx)
	return &newBankTx, nil
}
func (db *MongoDB) DeleteBankTx(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectBankTx()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetBankTxs(ctx context.Context, filterObj *account.BankTx, page *db.Pagination) ([]*account.BankTx, error) {
	client, coll := db.ConnectBankTx()
	defer MongoDisconnect(client)
	option := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}})
	if !reflect.ValueOf(page).IsZero() {
		if page.Skip != 0 {
			option.SetSkip(page.Skip)
		}
		if page.Limit != 0 {
			option.SetLimit(page.Limit)
		}
	}
	var filter = bson.M{}
	if !reflect.ValueOf(filterObj).IsZero() {
		filterItems := []bson.M{}
		if filterObj.XId != "" {
			id, err := primitive.ObjectIDFromHex(filterObj.XId)
			if err != nil {
				return nil, nil
			}
			filterItems = append(filterItems, bson.M{"_id": id})
		}

		if filterObj.BankId != "" {
			filterItems = append(filterItems, bson.M{"bank_id": filterObj.BankId})
		}
		if filterObj.PaymentType != "" {
			filterItems = append(filterItems, bson.M{"payment_type": filterObj.PaymentType})
		}
		if filterObj.EmployeeId != "" {
			filterItems = append(filterItems, bson.M{"employee_id": filterObj.EmployeeId})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		bankTxs    []*account.BankTx
		jsonBankTx []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonBankTx); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonBankTx, &bankTxs)
	return bankTxs, nil
}
func (db *MongoDB) CreateBankTx(ctx context.Context, bankTx *account.BankTx) (*account.BankTx, error) {
	client, coll := db.ConnectBankTx()
	defer MongoDisconnect(client)
	var (
		jsonBankTx interface{}
		result     bson.M
		newBankTx  account.BankTx
	)
	common.ToJSONStruct(bankTx, &jsonBankTx)
	insertResult, err := coll.InsertOne(ctx, jsonBankTx)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newBankTx)
	return &newBankTx, nil
}
