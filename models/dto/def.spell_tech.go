package dto

type CreateSpellSchoolRequest struct {
	School string `json:"school" form:"school" binding:"required"`
	Cost string `json:"cost" form:"cost" binding:""`
}

type CreateTechTypeRequest struct {
	Tech string `json:"tech" form:"tech" binding:"required"`
	Fuel string `json:"fuel" form:"fuel" binding:""`
}

type EditSpellSchoolBase struct {
	School string `json:"school" form:"school" binding:"required"`
	Cost string `json:"cost" form:"cost" binding:""`
}

type EditTechTypeBase struct {
	Tech string `json:"tech" form:"tech" binding:"required"`
	Fuel string `json:"fuel" form:"fuel" binding:""`
}

type AddAbilityToSpellSchool struct {
	Ability string `json:"ability" form:"ability" binding:"required"`
	School string `json:"school" form:"school" binding:"required"`
}

type AddAbilityToTechType struct {
	Ability string `json:"ability" form:"ability" binding:"required"`
	Tech string `json:"tech" form:"tech" binding:"required"`
}

type RemoveAbilityFromSpellSchool struct {
	Ability string `json:"ability" form:"ability" binding:"required"`
	School string `json:"school" form:"school" binding:"required"`
}

type RemoveAbilityFromTechType struct {
	Ability string `json:"ability" form:"ability" binding:"required"`
	Tech string `json:"tech" form:"tech" binding:"required"`
}

type EditAbilityOnSpellSchool struct {
	Ability string `json:"ability" form:"ability" binding:"required"`
	NewValue string `json:"newvalue" form:"newvalue" binding:"required"`
	School string `json:"school" form:"school" binding:"required"`
}

type EditAbilityOnTechType struct {
	Ability string `json:"ability" form:"ability" binding:"required"`
	NewValue string `json:"newvalue" form:"newvalue" binding:"required"`
	Tech string `json:"tech" form:"tech" binding:"required"`
}