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
	// WARNING: before running the following commands, you must first activate a mongod in your local terminal <mongod --auth --port 27017 --dbpath /usr/local/var/mongodb>
	// WARNING: you must first create a user who has access to the db specified by MONGO_DB_NAME via your local mongodb's superadmin
	// establish connection
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	authSource := os.Getenv("MONGO_AUTH_SOURCE")
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	credential := options.Credential{
		AuthSource: authSource,
		Username: username,
		Password: password,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		panic(err)
	}
	// check connection
	ctxPing, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
        panic(err)
	}
	// bind client to global instance variables
	MongoClient = client
	MongoDB = client.Database(dbName)
	MongoContext = ctx
}

// this function should rarely be used. The only times when it should be used is to check client info
func GetMongoClient() *mongo.Client {
	return MongoClient
}

// use this to perform crud
func GetMongoDB() *mongo.Database{
	return MongoDB
}

func CloseMongoDBConnection(db *mongo.Client) {
	db.Disconnect(MongoContext)
}