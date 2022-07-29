package entities

import "gorm.io/gorm"

type UserExtras struct{
	gorm.Model
	Gender string `gorm:"column:gender"`
	Organization Organization `gorm:"foreignKey:Organization"`
	User User `gorm:"foreignKey:Username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Organization struct {
	OrganizationGroup string `gorm:"column:organization_group"`
	Organization string `gorm:"column:organization;primaryKey"`
	Department string `gorm:"column:department"`
	Role string `gorm:"column:role"`
}