package migration

import (
	"github.com/dasha-kinsely/ostruct/models/entities"
	"gorm.io/gorm"
)

func MigrateSqlDB(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.UserExtras{})
}