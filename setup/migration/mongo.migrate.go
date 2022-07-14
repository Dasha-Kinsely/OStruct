package migration

import (
	"context"
	"log"

	"github.com/dasha-kinsely/ostruct/setup/config"
	"go.mongodb.org/mongo-driver/bson"
)

func MigrateMongoDB() {
	userCollection := config.GetMongoClient().Database("local").Collection("users")
	user := bson.D{{"name", "u1"}, {"age", 40}}
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	log.Println(result)
}