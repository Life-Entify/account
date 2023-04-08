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
	PAYMENT_CAT_COLLECTION = "payment_categories"
)

func (db *MongoDB) ConnectPaymentCategory() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(PAYMENT_CAT_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdatePaymentCategory(ctx context.Context, _id primitive.ObjectID, cheque *account.PaymentCategory) (*account.PaymentCategory, error) {
	client, coll := db.ConnectPaymentCategory()
	defer MongoDisconnect(client)

	var jsonPaymentCategory bson.M
	common.ToJSONStruct(cheque, &jsonPaymentCategory)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonPaymentCategory})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result             bson.M
		newPaymentCategory account.PaymentCategory
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newPaymentCategory)
	return &newPaymentCategory, nil
}
func (db *MongoDB) DeletePaymentCategory(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectPaymentCategory()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetPaymentCategories(ctx context.Context, filterObj *account.PaymentCategory, page *db.Pagination) ([]*account.PaymentCategory, error) {
	client, coll := db.ConnectPaymentCategory()
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

		if filterObj.Type != "" {
			filterItems = append(filterItems, bson.M{"type": filterObj.Type})
		}
		if filterObj.Title != "" {
			filterItems = append(filterItems, bson.M{"title": filterObj.Title})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		cheques             []*account.PaymentCategory
		jsonPaymentCategory []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonPaymentCategory); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonPaymentCategory, &cheques)
	return cheques, nil
}
func (db *MongoDB) CreatePaymentCategory(ctx context.Context, cheque *account.PaymentCategory) (*account.PaymentCategory, error) {
	client, coll := db.ConnectPaymentCategory()
	defer MongoDisconnect(client)
	var (
		jsonPaymentCategory interface{}
		result              bson.M
		newPaymentCategory  account.PaymentCategory
	)
	common.ToJSONStruct(cheque, &jsonPaymentCategory)
	insertResult, err := coll.InsertOne(ctx, jsonPaymentCategory)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newPaymentCategory)
	return &newPaymentCategory, nil
}
