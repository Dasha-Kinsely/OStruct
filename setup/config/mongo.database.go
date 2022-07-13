package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoDB *mongo.Database
	MongoClient *mongo.Client
	Context *context.Context
)

func SetupMongoDBConnection() {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := db.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(err)
	}
}

