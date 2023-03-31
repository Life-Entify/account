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
	CHEQUE_COLLECTION = "cheques"
)

func (db *MongoDB) ConnectCheque() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(CHEQUE_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdateCheque(ctx context.Context, _id primitive.ObjectID, cheque *account.Cheque) (*account.Cheque, error) {
	client, coll := db.ConnectCheque()
	defer MongoDisconnect(client)

	var jsonCheque bson.M
	common.ToJSONStruct(cheque, &jsonCheque)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonCheque})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result    bson.M
		newCheque account.Cheque
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newCheque)
	return &newCheque, nil
}
func (db *MongoDB) DeleteCheque(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectCheque()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetCheques(ctx context.Context, filterObj *account.Cheque, page *db.Pagination) ([]*account.Cheque, error) {
	client, coll := db.ConnectCheque()
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
		if filterObj.ChequeNumber != "" {
			filterItems = append(filterItems, bson.M{"cheque_number": filterObj.ChequeNumber})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		cheques    []*account.Cheque
		jsonCheque []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonCheque); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonCheque, &cheques)
	return cheques, nil
}
func (db *MongoDB) CreateCheque(ctx context.Context, cheque *account.Cheque) (*account.Cheque, error) {
	client, coll := db.ConnectCheque()
	defer MongoDisconnect(client)
	var (
		jsonCheque interface{}
		result     bson.M
		newCheque  account.Cheque
	)
	common.ToJSONStruct(cheque, &jsonCheque)
	insertResult, err := coll.InsertOne(ctx, jsonCheque)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newCheque)
	return &newCheque, nil
}
