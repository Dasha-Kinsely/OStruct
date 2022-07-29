package entities

type Ability struct {
	AbilitySpec string `gorm:"column:spec;notNull;primaryKey"`
}

// to be used as polymorphic component only
type Weapon struct {
	Weapon string
	HolderID uint
	HolderType string
}

type UnitDesign struct {
	
	Weapon Weapon `gorm:"polymorphic:Holder;"`
}