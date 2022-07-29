package entities

import "gorm.io/gorm"

// One-to-Many, polymorphic
type SpellSchool struct {
	gorm.Model
	School string `gorm:"column:school;unique;notNull"`
	Cost string `gorm:"column:cost"`
	AbilityScope []Ability `gorm:"foreignKey:AbilitySpec"`
	FullCt int64 `gorm:"column:full_ct"` // updated by triggers from mongodb
	PartialCt int64 `gorm:"column:partial_ct"` // updated by triggers from mongodb
	ItemCt int64 `gorm:"column:item_ct"` // updated by triggers from mongodb
}