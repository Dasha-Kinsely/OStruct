package entities

import "gorm.io/gorm"

// <Smart-Field-Obj>
type CC struct {
	gorm.Model
	ControlType string `gorm:"column:control_type;unique;primaryKey"`
	Classification string `gorm:"column:classification"` //has enum
	Movement string `gorm:"column:movement"` // has enum
	BasicAttack string `gorm:"column:basic_attack"` // has enum
	AbilityCast string `gorm:"column:ability_cast"` // has enum
	BuffEffects string `gorm:"column:buff_effects"`
	StatEffects string `gorm:"column:stat_effects"`
	MentalityMitigation bool `gorm:"column:mentality"`
	TenacityMitigation bool `gorm:"column:tenacity"`
	WeightMitigation bool `gorm:"column:weight"`
}