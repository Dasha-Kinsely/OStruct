package dto

type CreateNeutralMonsterRequest struct {
	MonsterType string `json:"monstertype" form:"monstertype" binding:"required"`
	Tankiness string `json:"tankiness" form:"tankiness" binding:"alphanum"`
	AttackPattern string `json:"attackpattern" form:"attackpattern" binding:""`
	PatrolPattern string `json:"patrolpattern" form:"patrolpattern" binding:""`
	Respawn string `json:"respawn" form:"respawn" binding:""`
	Reward string `json:"reward" form:"reward" binding:""`
}

type UpdateOneNeutralMonster struct {
	MonsterType string `json:"monstertype" form:"monstertype" binding:"required"`
	Tankiness string `json:"tankiness" form:"tankiness" binding:"alphanum"`
	AttackPattern string `json:"attackpattern" form:"attackpattern" binding:""`
	PatrolPattern string `json:"patrolpattern" form:"patrolpattern" binding:""`
	Respawn string `json:"respawn" form:"respawn" binding:""`
	Reward string `json:"reward" form:"reward" binding:""`
}

// ---------------------------------------------------------------------------------

type CreateSubordinateMonsterRequest struct {
	MonsterType string `json:"monstertype" form:"monstertype" binding:"required"`
	Tankiness string `json:"tankiness" form:"tankiness" binding:""`
	AttackPattern string `json:"attackpattern" form:"attackpattern" binding:""`
	Reward string `json:"reward" form:"reward" binding:""`
}

type UpdateOneSubordinateMonster struct {
	MonsterType string `json:"monstertype" form:"monstertype" binding:"required"`
	Tankiness string `json:"tankiness" form:"tankiness" binding:""`
	AttackPattern string `json:"attackpattern" form:"attackpattern" binding:""`
	Reward string `json:"reward" form:"reward" binding:""`
}