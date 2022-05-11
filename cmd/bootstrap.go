package cmd

import (
	"github.com/arymaulanamalik/sicepat_sample/pkg/database"
	"github.com/arymaulanamalik/sicepat_sample/pkg/logger"
)

var (
	mdb       *database.MongoDatabase
	configURL string
)

func initModule() {
	mdb = InitMongoModule()
	logger.Configure()
}

func InitMongoModule() *database.MongoDatabase {
	return database.NewMongoClient()
}
