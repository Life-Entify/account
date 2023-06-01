package account

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	common "github.com/life-entify/account/common"
	errors "github.com/life-entify/account/errors"
	repo_db "github.com/life-entify/account/repository/db"
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
func paymentFilter(paymentFilter *account.Payment) (bson.M, error) {
	var filter = bson.M{}
	if !reflect.ValueOf(paymentFilter).IsZero() {
		filterItems := []bson.M{}
		if paymentFilter.XId != "" {
			id, err := primitive.ObjectIDFromHex(paymentFilter.XId)
			if err != nil {
				return nil, errors.Errorf(err.Error())
			}
			filterItems = append(filterItems, bson.M{"_id": id})
		}

		if paymentFilter.BankId != "" {
			filterItems = append(filterItems, bson.M{"bank_id": paymentFilter.BankId})
		}
		if paymentFilter.PersonId != 0 {
			filterItems = append(filterItems, bson.M{"person_id": paymentFilter.PersonId})
		}
		if paymentFilter.ActionType != "" {
			filterItems = append(filterItems, bson.M{"action_type": paymentFilter.ActionType})
		}
		if paymentFilter.PayType != "" {
			filterItems = append(filterItems, bson.M{"pay_type": paymentFilter.PayType})
		}
		if paymentFilter.TxType != "" {
			filterItems = append(filterItems, bson.M{"tx_type": paymentFilter.TxType})
		}
		if paymentFilter.Unresolved != "" {
			filterItems = append(filterItems, bson.M{"unresolved": paymentFilter.Unresolved})
		}
		if paymentFilter.EmployeeId != 0 {
			filterItems = append(filterItems, bson.M{"employee_id": paymentFilter.EmployeeId})
		}
		if len(filterItems) > 0 {
			filter["$or"] = filterItems
		}
	}
	return filter, nil
}
func paymentDateFilter(dateFilter *repo_db.DateFilter) ([]bson.D, error) {
	if dateFilter != nil {
		minInt, _ := strconv.Atoi(dateFilter.DateStampFrom)
		maxInt, _ := strconv.Atoi(dateFilter.DateStampTo)
		if maxInt == 0 {
			maxInt = minInt
		}
		filter := bson.D{{Key: "$gte", Value: minInt}}
		if minInt != 0 {
			filter = append(filter, bson.E{Key: "$lte", Value: maxInt})
		}
		return []bson.D{
			{{Key: "$addFields", Value: bson.D{{Key: "createdAtInt", Value: bson.D{{Key: "$toLong", Value: "$created_at"}}}}}},
			{{Key: "$match",
				Value: bson.D{{
					Key:   "createdAtInt",
					Value: filter,
				}},
			}},
		}, nil
	}
	return nil, fmt.Errorf("dateFilter can not be nil if used")
}
func (db *MongoDB) GetDepositByPersonGroup(ctx context.Context, page *repo_db.Pagination) ([]*repo_db.DepositSummary, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{
			Key: "action_type", Value: bson.D{{
				Key: "$in", Value: []string{"receive_deposit", "use_deposit", "deposit_withdrawal"}}},
		}}}},
	}
	if !reflect.ValueOf(page).IsZero() {
		if page.Skip != 0 {
			pipeline = append(pipeline, bson.D{{
				Key: "$skip", Value: page.Skip,
			}})
		}
		if page.Limit != 0 {
			pipeline = append(pipeline, bson.D{{
				Key: "$limit", Value: page.Limit,
			}})
		}
	}
	pipeline = append(pipeline, bson.D{{Key: "$group", Value: bson.M{
		"_id": bson.D{
			{Key: "person_id", Value: "$person_id"},
			{Key: "action_type", Value: "$action_type"},
		},
		"total_amount": bson.M{"$sum": "$total_amount"},
	}}})
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(context.Background())
	var (
		results    []*repo_db.DepositSummary
		jsonResult []*bson.M
	)
	if err = cursor.All(context.Background(), &jsonResult); err != nil {
		return nil, errors.Errorf(err.Error())
	}
	common.ToJSONStruct(jsonResult, &results)
	return results, nil
}

func (db *MongoDB) GetPaymentByEmpIdAndPayType(ctx context.Context, filterObj *account.Payment, dateFilter *repo_db.DateFilter) ([]*repo_db.PaymentTypeSummary, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	filter, _ := paymentFilter(filterObj)
	// Set up pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
	}
	if dateFilter != nil {
		dateMatch, err := paymentDateFilter(dateFilter)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		pipeline = append(pipeline, dateMatch...)
	}
	pipeline = append(pipeline, bson.D{{Key: "$group", Value: bson.M{
		"_id": bson.D{
			{Key: "employee_id", Value: "$employee_id"},
			{Key: "pay_type", Value: "$pay_type"},
			{Key: "tx_type", Value: "$tx_type"},
		},
		"total_amount": bson.M{"$sum": "$total_amount"},
	}}})
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(context.Background())
	var (
		results    []*repo_db.PaymentTypeSummary
		jsonResult []*bson.M
	)
	if err = cursor.All(context.Background(), &jsonResult); err != nil {
		return nil, errors.Errorf(err.Error())
	}

	common.ToJSONStruct(jsonResult, &results)
	return results, nil
}
func (db *MongoDB) GetPaymentByEmpIdAndActionType(ctx context.Context, filterObj *account.Payment, dateFilter *repo_db.DateFilter) ([]*repo_db.PaymentActionTypeSummary, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	filter, _ := paymentFilter(filterObj)
	// Set up pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
	}
	if dateFilter != nil {
		dateMatch, err := paymentDateFilter(dateFilter)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		pipeline = append(pipeline, dateMatch...)
	}
	pipeline = append(pipeline, bson.D{{Key: "$group", Value: bson.M{
		"_id": bson.D{
			{Key: "employee_id", Value: "$employee_id"},
			{Key: "action_type", Value: "$action_type"},
		},
		"total_amount": bson.M{"$sum": "$total_amount"},
	}}})
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(context.Background())
	var (
		results    []*repo_db.PaymentActionTypeSummary
		jsonResult []*bson.M
	)
	if err = cursor.All(context.Background(), &jsonResult); err != nil {
		return nil, errors.Errorf(err.Error())
	}

	common.ToJSONStruct(jsonResult, &results)
	return results, nil
}
func (db *MongoDB) GetPaymentEmployeeIds(ctx context.Context, filterObj *account.Payment) ([]int64, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	filter, _ := paymentFilter(filterObj)
	result, err := coll.Distinct(ctx, "employee_id", filter)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var employeeIds []int64
	common.ToJSONStruct(result, &employeeIds)
	return employeeIds, nil
}
func (db *MongoDB) GetPayments(ctx context.Context, filterObj *account.Payment, dateFilter *repo_db.DateFilter, page *repo_db.Pagination) ([]*account.Payment, error) {
	client, coll := db.ConnectPayment()
	defer MongoDisconnect(client)
	pipeline := mongo.Pipeline{
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: -1}}}},
	}
	if !reflect.ValueOf(page).IsZero() {
		if page.Skip != 0 {
			pipeline = append(pipeline, bson.D{{
				Key: "$skip", Value: page.Skip,
			}})
		}
		if page.Limit != 0 {
			pipeline = append(pipeline, bson.D{{
				Key: "$limit", Value: page.Limit,
			}})
		}
	}
	filter, _ := paymentFilter(filterObj)
	if filter != nil {
		pipeline = append(pipeline, bson.D{{
			Key: "$match", Value: filter,
		}})
	}
	if dateFilter != nil {
		dateMatch, err := paymentDateFilter(dateFilter)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		pipeline = append(pipeline, dateMatch...)
	}
	var (
		payments    []*account.Payment
		jsonPayment []*bson.M
	)
	cursor, err := coll.Aggregate(ctx, pipeline)
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
