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
	PAYROLL_ACTION_COLLECTION = "payroll_actions"
)

func (db *MongoDB) ConnectPayrollAction() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(PAYROLL_ACTION_COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func (db *MongoDB) UpdatePayrollAction(ctx context.Context, _id primitive.ObjectID, payrollAction *account.PayrollAction) (*account.PayrollAction, error) {
	client, coll := db.ConnectPayrollAction()
	defer MongoDisconnect(client)

	var jsonPayrollAction bson.M
	common.ToJSONStruct(payrollAction, &jsonPayrollAction)
	_, err := coll.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": jsonPayrollAction})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var (
		result           bson.M
		newPayrollAction account.PayrollAction
	)
	coll.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newPayrollAction)
	return &newPayrollAction, nil
}
func (db *MongoDB) DeletePayrollAction(ctx context.Context, _id primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.ConnectPayrollAction()
	defer MongoDisconnect(client)
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return deleteResult, nil
}
func (db *MongoDB) GetPayrollActions(ctx context.Context, filterObj *account.PayrollAction, page *db.Pagination) ([]*account.PayrollAction, error) {
	client, coll := db.ConnectPayrollAction()
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

		if filterObj.ActionKind != "" {
			filterItems = append(filterItems, bson.M{"action_kin": filterObj.ActionKind})
		}
		if filterObj.ActionType != "" {
			filterItems = append(filterItems, bson.M{"action_type": filterObj.ActionType})
		}
		if filterObj.Active != "" {
			filterItems = append(filterItems, bson.M{"active": filterObj.Active})
		}
		if filterObj.IsConstant != "" {
			filterItems = append(filterItems, bson.M{"is_constant": filterObj.IsConstant})
		}
		if filterObj.Name != "" {
			filterItems = append(filterItems, bson.M{"name": filterObj.Name})
		}
		if filterObj.IsGeneral != "" {
			filterItems = append(filterItems, bson.M{"is_general": filterObj.IsGeneral})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	var (
		payrollActions    []*account.PayrollAction
		jsonPayrollAction []*bson.M
	)
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &jsonPayrollAction); err != nil {
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
	}
	common.ToJSONStruct(jsonPayrollAction, &payrollActions)
	return payrollActions, nil
}
func (db *MongoDB) CreatePayrollAction(ctx context.Context, payrollAction *account.PayrollAction) (*account.PayrollAction, error) {
	client, coll := db.ConnectPayrollAction()
	defer MongoDisconnect(client)
	var (
		jsonPayrollAction interface{}
		result            bson.M
		newPayrollAction  account.PayrollAction
	)
	common.ToJSONStruct(payrollAction, &jsonPayrollAction)
	insertResult, err := coll.InsertOne(ctx, jsonPayrollAction)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	err = coll.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(result, &newPayrollAction)
	return &newPayrollAction, nil
}
