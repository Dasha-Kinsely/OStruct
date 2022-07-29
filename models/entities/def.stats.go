package entities

import "gorm.io/gorm"

type Stats struct {
	gorm.Model
	Defensive uint `gorm:"column:defensive"` 
	DefensiveRef DefensiveStats `gorm:"foreignKey:ID;references:Defensive;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Utility uint `gorm:"column:utility"`
	UtilityRef UtilityStats `gorm:"foreignKey:ID;references:Utility;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	BasicAttack uint `gorm:"column:basic_attack"`
	BasicAttackRef BasicAttackStats `gorm:"foreignKey:ID;references:BasicAttack;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

type DefensiveStats struct {
	gorm.Model
	HP int64 `gorm:"column:hp"`
	SelfRegeneration float32 `gorm:"column:self_regeneration"`
	Effectiveness float32 `gorm:"column:effectiveness"`// healing and shielding
	MythicResistance int64 `gorm:"column:mythic_resistance"`
	PhysicalResistance int64 `gorm:"column:physical_resistance"`
	MaterialResistance int64 `gorm:"column:material_resistance"`// equals to 80% of the avg between MR and PR
	Blocking int64 `gorm:"column:blocking"`
	NetDmgReduction float32 `gorm:"column:net_dmg_reduction"`
}

type UtilityStats struct {
	gorm.Model
	MovementSpeed int64 `gorm:"column:movement_speed"`
	TrekSpeed int64 `gorm:"column:trek_speed"`
	Tenacity float32 `gorm:"column:tenacity"`
	Mentality float32 `gorm:"column:mentality"`
	Weight int64 `gorm:"column:weight"`
	SightRange int64 `gorm:"column:sight_range"`
	HitboxSize int64 `gorm:"column:hitbox_size"`
	CooldownReduction float32 `gorm:"column:cooldown_reduction"`
}

type BasicAttackStats struct {
	gorm.Model
	AttackRecovery float32 `gorm:"column:attack_recovery"`
	AttackPrecast float32 `gorm:"column:attack_precast"`
}

//----------------------------------------------------------------------
type MageStats struct {
	gorm.Model
	MythicPower int64 `gorm:"column:mythic_power"`
	MythicalPenetration int64 `gorm:"column:mythic_penetration"`
	MythicPenetrationPercent float32 `gorm:"column:mythic_penetration_%"`
	Receptivity int64 `gorm:"column:receptivity"`
	LifeDrain float32 `gorm:"column:life_drain"`
	Mana int64 `gorm:"column:mana"`
	ManaRegeneration int64 `gorm:"column:mana_regeneration"`
}

type MaterialStats struct {
	gorm.Model
	Potency int64 `gorm:"column:potency"`
}

type MechanicalStats struct {
	gorm.Model
	Mechanization int64 `gorm:"column:mechanization"`
	Penetration int64 `gorm:"column:penetration"`
}

type PhysicalStats struct {
	gorm.Model
	Strength int64 `gorm:"column:strength"`
	Penetration int64 `gorm:"column:penetration"`
	PenetrationPercent float32 `gorm:"column:penetration_%"`
	LifeVamp float32 `gorm:"column:life_vamp"`
}