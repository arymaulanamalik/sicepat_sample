package database

import (
	"sync"
	"time"

	mongolib "gitlab.sicepat.tech/platform/golib/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabase struct {
	DB *mongo.Database
}

var (
	MongoClient *MongoDatabase
	hasOnce     sync.Once
)

func NewMongoClient() *MongoDatabase {
	hasOnce.Do(func() {
		db := mongolib.MongoConnectClient(&mongolib.Client{
			URI:            "mongodb://localhost:27017",
			DB:             "authorizationdb",
			AppName:        "authorization-svc",
			ConnectTimeout: time.Duration(30) * time.Second,
			PingTimeout:    time.Duration(10) * time.Second,
		})
		MongoClient = &MongoDatabase{DB: db.Database}
	})
	return MongoClient
}
