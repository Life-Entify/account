package account

import (
	config "github.com/life-entify/account/config"
	repo "github.com/life-entify/account/repository"
	db "github.com/life-entify/account/repository/db/mongo"
)

type Controller struct {
	repo.Repository
	Config *config.Config
}

const (
	Mongo    = "MONGODB"
	MySQL    = "MYSQL"
	PostGres = "POSTGRES"
)

func NewController(config *config.Config) *Controller {
	return &Controller{
		db.NewMongoDB(config),
		config,
	}
}
