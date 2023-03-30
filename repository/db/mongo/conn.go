package account

import (
	"context"

	config "github.com/life-entify/account/config"
	"go.mongodb.org/mongo-driver/mongo"
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
func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
