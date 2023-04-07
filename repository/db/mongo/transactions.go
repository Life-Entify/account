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
	TX_COLLECTION = "transactions"
)

func (db *MongoDB) ConnectTransaction() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(TX_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdateTransaction(ctx context.Context, _id primitive.ObjectID, transaction *account.Transaction) (*account.Transaction, error) {
	client, coll := db.ConnectTransaction()
	defer MongoDisconnect(client)

	var jsonTransaction bson.M
	common.ToJSONStruct(transaction, &jsonTransaction)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonTransaction})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result         bson.M
		newTransaction account.Transaction
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newTransaction)
	return &newTransaction, nil
}
func (db *MongoDB) DeleteTransaction(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectTransaction()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetTransactions(ctx context.Context, filterObj *account.Transaction, page *db.Pagination) ([]*account.Transaction, error) {
	client, coll := db.ConnectTransaction()
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

		if filterObj.PaymentId != 0 {
			filterItems = append(filterItems, bson.M{"payment_id": filterObj.PaymentId})
		}
		if filterObj.CategoryId != "" {
			filterItems = append(filterItems, bson.M{"category_id": filterObj.CategoryId})
		}
		if filterObj.TxType != "" {
			filterItems = append(filterItems, bson.M{"tx_type": filterObj.TxType})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		transactions    []*account.Transaction
		jsonTransaction []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonTransaction); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonTransaction, &transactions)
	return transactions, nil
}
func (db *MongoDB) CreateTransaction(ctx context.Context, transaction *account.Transaction) (*account.Transaction, error) {
	client, coll := db.ConnectTransaction()
	defer MongoDisconnect(client)
	var (
		jsonTransaction interface{}
		result          bson.M
		newTransaction  account.Transaction
	)
	common.ToJSONStruct(transaction, &jsonTransaction)
	insertResult, err := coll.InsertOne(ctx, jsonTransaction)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newTransaction)
	return &newTransaction, nil
}
