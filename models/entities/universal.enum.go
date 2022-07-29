package entities

type ControlClassificationEnum int8
const (
	Hard ControlClassificationEnum=iota+1
	Medium
	Soft
	Specialized
)
func (cont ControlClassificationEnum) ToString() string {
	return [...]string{"Hard", "Medium", "Soft", "Specialized"}[cont-1]
} 

type MovementEnum int8
const (
	MovControlled MovementEnum=iota+1
	MovForbidden
	Limited
	Reduced
	Dimensional
	Undashable
	Stopped
	Morphed
)
func (mov MovementEnum) ToString() string{
	return [...]string{"Controlled", "Forbidden", "Limited", "Reduced", "Dimensional", "Undashable", "Stopped", "Morphed"}[mov-1]
}

type BasicAttackEnum int8
const (
	AtkControlled BasicAttackEnum=iota+1
	AtkForbidden 
	AtkReplaced
	MagForbidden
	AtkHalted
	AtkTeamSwapped
	AtkDelayed
	AtkGrounded
	AtkDimensional
)
// TODO: write these
func (atk BasicAttackEnum) ToString() string {
	return [...]string{}[atk-1]
}

type AbilityEnum int8
const (
	CastForbidden AbilityEnum=iota+1
	CastCanceled
	CastTeamSwapped
	CastDelayed
	CastGrounded
	CastDimensional
)
// TODO: write these
func (cast AbilityEnum) ToString() string {
	return [...]string{}[cast-1]
}

type EquipmentEnum int8
const (
	LargeWearable EquipmentEnum=iota+1
	SmallWearable 
	MainWeapon
	OffWeapon
	Enchantment 
	BioAugment
)

func (eqp EquipmentEnum) ToString() string {
	return [...]string{"Large Wearable", "Small Wearable", "Main Weapon", "Offhand Weapon", "Enchantment", "Bio-Augment"}[eqp-1]
}

type AbundanceEnum int8
const (
	Everywhere AbundanceEnum=iota+1
	Common
	Sparse
	Rare
	Unique
)
func (abundance AbundanceEnum) ToString() string {
	return [...]string{"Everywhere", "Common", "Sparse", "Rare", "Unique"}[abundance-1]
}

type IntelligenceEnum int8
const (
	High IntelligenceEnum=iota+1
	AboveAverage
	Fluctuating
	Average
	Low
	Beastly
)
func (intel IntelligenceEnum) ToString() string{
	return [...]string{"High", "Above Average", "Fluctuating", "Average", "Low", "Beastly"}[intel-1]
}

// TODO:.........
type GenderEnum int8 
const (
	Male GenderEnum=iota+1
	Female
	Other
)
func (gender GenderEnum) ToString() string {
	return [...]string{"Male", "Female", "Other"}[gender-1]
}

type MajorCategoryEnum int8
const (
	Mage MajorCategoryEnum=iota+1
	NonMage
	Spectral
)
func (cate MajorCategoryEnum) ToString() string{
	return [...]string{"Mage", "NonMage", "Spectral"}[cate-1]
}

type DmgTypeEnum int8
const (
	Magic DmgTypeEnum=iota+1
	Material
	Mechanical
	Strength
)
func (dt DmgTypeEnum) ToString() string{
	return [...]string{"Magic", "Material", "Mechanical", "Strength"}[dt-1]
}

type ClassEnum int8
const (
	Tech ClassEnum=iota+1
	Spell
	Phys
) 
func (class ClassEnum) ToString() string{
	return [...]string{"Technological", "Spell", "Physical"}[class-1]
}
