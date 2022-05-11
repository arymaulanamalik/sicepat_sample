package cmd

import (
	"github.com/maul/sicepat_sample/pkg/database"
	"github.com/maul/sicepat_sample/pkg/logger"
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
