package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoDB *mongo.Database
	MongoClient *mongo.Client
	MongoContext context.Context
)

func SetupMongoDBConnection() {
	uri := os.Getenv("MONGO_URI")
	dbname := os.Getenv("MONGO_DB_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctxPing, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
        panic(err)
	}
	MongoClient = client
	MongoDB = client.Database(dbname)
	MongoContext = ctx
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}

func GetMongoDB() *mongo.Database{
	return MongoDB
}

func CloseMongoDBConnection(db *mongo.Client) {
	db.Disconnect(MongoContext)
}