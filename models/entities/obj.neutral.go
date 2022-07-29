package entities

import "gorm.io/gorm"

// One-to-One
type Monster struct {
	gorm.Model
	MonsterType string `gorm:"column:monster_type;unique"`
	Tankiness int64 `gorm:"column:tankiness"`
	AttackPattern string `gorm:"column:attack_pattern"`
	PatrolPattern string `gorm:"column:patrol_pattern"`
	Respawn string `gorm:"column:respawn"`
	Reward string `gorm:"column:reward"`
	// -------------------- Owns "Subordinate" -----------------------------------------
	SubordinateRef string `gorm:"column:subordinate_ref"`
	Subordinates Subordinate `gorm:"foreignKey:MonsterType;references:SubordinateRef;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	SubordinateCt int64 `gorm:"column:subordinate_count"`
}

type Subordinate struct {
	gorm.Model
	MonsterType string `gorm:"column:monster_type;unique;primaryKey"`
	Tankiness int64 `gorm:"column:tankiness"`
	AttackPattern string `gorm:"column:attack_pattern"`
	Reward string `gorm:"column:reward"`
}