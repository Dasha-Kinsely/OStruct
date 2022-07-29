package entities

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Class string `gorm:"column:class"`
	Category string `gorm:"column:category"`
	DmgType string `gorm:"column:dmg_type"`
	// This identifies which obj to use, usually filtered by validators based on major-category and dmg-type auto-generated
	IsA string `gorm:"column:is_a"`
}
// ---------------------------------
type Technology struct {
	gorm.Model
	TechnologyType string
	Styles []PotentialStyle
}

type School struct {
	gorm.Model
	MagicType string
	Styles []PotentialStyle
}

type Physical struct {
	gorm.Model
	BodyType string
	Styles []PotentialStyle
}

type PotentialStyle struct {
	gorm.Model
	Weapon string
	Style string
	Perks string
	PreRequisite string
}