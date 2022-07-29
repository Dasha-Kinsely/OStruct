package migration

import (
	"github.com/dasha-kinsely/ostruct/models/entities"
	"gorm.io/gorm"
)

// WARNING: the order matters when migrating relationships with foreign key constriants, always migrate the master tables after their slave tables
func MigrateSqlDB(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.UserExtras{})
	//-------------------------------------
	db.AutoMigrate(&entities.Head{})
	db.AutoMigrate(&entities.Body{})
	db.AutoMigrate(&entities.Menance{})
	db.AutoMigrate(&entities.Race{})
	//-------------------------------------
	db.AutoMigrate(&entities.Subordinate{})
	db.AutoMigrate(&entities.Monster{})
	//-------------------------------------
	db.AutoMigrate(&entities.Equipment{})
	//-------------------------------------
	db.AutoMigrate(&entities.CC{})
	//-------------------------------------
	db.AutoMigrate(&entities.SpellSchool{})
	db.AutoMigrate(&entities.TechType{})
}