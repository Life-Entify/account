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
	PAYMENT_COLLECTION = "payments"
)

func (db *MongoDB) ConnectPayment() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(PAYMENT_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdatePayment(ctx context.Context, _id primitive.ObjectID, payment *account.Payment) (*account.Payment, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)

	var jsonPayment bson.M
	common.ToJSONStruct(payment, &jsonPayment)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonPayment})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result     bson.M
		newPayment account.Payment
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newPayment)
	return &newPayment, nil
}
func (db *MongoDB) DeletePayment(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetPayments(ctx context.Context, filterObj *account.Payment, page *db.Pagination) ([]*account.Payment, error) {
	client, coll := db.ConnectPayment()
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
		if filterObj.PersonId != 0 {
			filterItems = append(filterItems, bson.M{"person_id": filterObj.PersonId})
		}
		if filterObj.PayType != "" {
			filterItems = append(filterItems, bson.M{"pay_type": filterObj.PayType})
		}
		if filterObj.TxType != "" {
			filterItems = append(filterItems, bson.M{"tx_type": filterObj.TxType})
		}
		if filterObj.Unresolved != "" {
			filterItems = append(filterItems, bson.M{"unresolved": filterObj.Unresolved})
		}
		if filterObj.EmployeeId != 0 {
			filterItems = append(filterItems, bson.M{"employee_id": filterObj.EmployeeId})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		payments    []*account.Payment
		jsonPayment []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonPayment); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonPayment, &payments)
	return payments, nil
}
func (db *MongoDB) GetPaymentByID(ctx context.Context, _id primitive.ObjectID) (*account.Payment, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)

	var (
		jsonPayment   bson.M
		resultPayment account.Payment
	)
	err := coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&jsonPayment)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(jsonPayment, &resultPayment)
	return &resultPayment, nil
}
func (db *MongoDB) CreatePayment(ctx context.Context, payment *account.Payment) (*account.Payment, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	var (
		jsonPayment interface{}
		result      bson.M
		newPayment  account.Payment
	)
	common.ToJSONStruct(payment, &jsonPayment)
	insertResult, err := coll.InsertOne(ctx, jsonPayment)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newPayment)
	return &newPayment, nil
}
