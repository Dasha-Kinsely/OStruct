package entities

import "gorm.io/gorm"

type Race struct {
	gorm.Model
	Race string `gorm:"column:race;primaryKey"`
	Family string `gorm:"column:family"`
	Head string `gorm:"column:head"`
	HeadRef Head `gorm:"foreignKey:ForRace;references:Head;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Body string `gorm:"column:body"`
	BodyRef Body `gorm:"foreignKey:ForRace;references:Body;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Menance string `gorm:"column:menance"`
	MenanceRef Menance `gorm:"foreignKey:ForRace;references:Menance;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

type Menance struct {
	gorm.Model
	ForRace string `gorm:"primaryKey"`
	Abundance string `gorm:"column:abundance"` // has enum
	Intelligence string `gorm:"column:intelligence"` // has enum
	MagicPrevalency int64 `gorm:"column:magic_prevalency"`
	Consolidarity string `gorm:"column:consolidarity"`
	AggressionIndex int64 `gorm:"column:aggression_index"`
	RacialPerks string `gorm:"column:racial_perks;size=2048"`
	MenanceScore int64 `gorm:"column:menance_score"`
}

type Head struct {
	gorm.Model
	ForRace string `gorm:"primaryKey"`
	Teeth string `gorm:"column:teeth"`
	Ears string `gorm:"column:ears"`
	Nose string `gorm:"column:nose"`
	Eyes string `gorm:"column:eyes"`
	Forehead string `gorm:"column:forehead"`
	Hair string `gorm:"column:hair"`
}

type Body struct {
	gorm.Model
	ForRace string `gorm:"primaryKey"`
	Height float32 `gorm:"column:height"`
	Muscle string `gorm:"column:muscle"`
	SkinHardness int64 `gorm:"column:skin_hardness"`
}