package config

import (
	"os"

	"github.com/dasha-kinsely/ostruct/setup/migration"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func BulkMigrate(db *gorm.DB, fdb *mongo.Client, cache *redis.Client) {
	//migration.MigrateMongoDB()
	migration.MigrateSqlDB(db)
}

func StartupCheck() bool {
	// does this run require setting up database schemas
	if os.Getenv("REQUIRES_MIGRATION") == "yes" {
		return true
	} else {
		return false
	}
}
