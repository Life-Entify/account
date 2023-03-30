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
	COLLECTION = "banks"
)

func (db *MongoDB) ConnectBank() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdateBank(ctx context.Context, _id primitive.ObjectID, bank *account.Bank) (*account.Bank, error) {
	client, coll := db.ConnectBank()
	defer MongoDisconnect(client)

	var jsonBank bson.M
	common.ToJSONStruct(bank, &jsonBank)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonBank})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result  bson.M
		newBank account.Bank
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newBank)
	return &newBank, nil
}
func (db *MongoDB) DeleteBank(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectBank()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetBanks(ctx context.Context, filterObj *account.Bank, page *db.Pagination) ([]*account.Bank, error) {
	client, coll := db.ConnectBank()
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

		if filterObj.Name != "" {
			filterItems = append(filterItems, bson.M{"name": filterObj.Name})
		}
		if filterObj.Number != "" {
			filterItems = append(filterItems, bson.M{"number": filterObj.Number})
		}
		if filterObj.IsAdmin != "" {
			filterItems = append(filterItems, bson.M{"is_admin": filterObj.IsAdmin})
		}
		if filterObj.Active != "" {
			filterItems = append(filterItems, bson.M{"active": filterObj.Active})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		banks    []*account.Bank
		jsonBank []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonBank); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonBank, &banks)
	return banks, nil
}
func (db *MongoDB) CreateBank(ctx context.Context, bank *account.Bank) (*account.Bank, error) {
	client, coll := db.ConnectBank()
	defer MongoDisconnect(client)
	var (
		jsonBank interface{}
		result   bson.M
		newBank  account.Bank
	)
	common.ToJSONStruct(bank, &jsonBank)
	insertResult, err := coll.InsertOne(ctx, jsonBank)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newBank)
	return &newBank, nil
}
