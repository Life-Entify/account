package account

import (
	"context"

	config "github.com/life-entify/person/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION = "persons"
)

type MongoDB struct {
	uri      string
	database string
}

func NewMongoDB(config config.IConfig) *MongoDB {
	return &MongoDB{
		uri:      config.GetDBUrl(),
		database: config.GetDBName(),
	}
}
func (db *MongoDB) Connect() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
