package entities

import "gorm.io/gorm"

// Create custom enum object
type GenderEnum int8 
const (
	Male GenderEnum=iota+1
	Female
	Other
)
func (gender GenderEnum) ToString() string {
	return [...]string{"Male", "Female", "Other"}[gender-1]
}

type UserExtras struct{
	gorm.Model
	Gender GenderEnum `gorm:"column:gender"`
	User User `gorm:"foreignKey:Username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}