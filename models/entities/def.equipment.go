package entities

import "gorm.io/gorm"

// <Smart-Field-Obj>
type Equipment struct {
	gorm.Model
	EquipmentType string `gorm:"column:equipment_type;unique"` 
	Category string `gorm:"column:category"` // has enum
	// filtered based on EquipmentEnum
	EquipLimit int64 `gorm:"column:equip_limit"`
}