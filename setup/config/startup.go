package config

import (
	"os"

	"github.com/dasha-kinsely/ostruct/setup/migration"
	"gorm.io/gorm"
)

func BulkMigrate(db *gorm.DB) {
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
