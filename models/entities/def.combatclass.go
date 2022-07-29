package entities

import "gorm.io/gorm"

// One-to-Many, polymorphic
type CombatClass struct {
	gorm.Model
	Class string `gorm:"column:class;unique;notNull"`
	AbilityScope []Ability `gorm:"foreignKey:AbilitySpec"`
	FullCt int64 `gorm:"column:full_ct"` // updated by triggers from mongodb
	PartialCt int64 `gorm:"column:partial_ct"` // updated by triggers from mongodb
	ItemCt int64 `gorm:"column:item_ct"` // updated by triggers from mongodb
}