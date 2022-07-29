package utils

import "github.com/dasha-kinsely/ostruct/models/entities"

func InterpretEnum(entity string, code int8) string {
	if code < 1 {
		return ""
	}
	switch entity {
	case "equipment":
		if code > 6 {
			return ""
		} else {
			entity = entities.EquipmentEnum(code).ToString()
		}
	case "abundance":
		if code > 5 {
			return ""
		} else {
			entity = entities.AbundanceEnum(code).ToString()
		}
	case "intelligence":
		if code > 6 {
			return ""
		} else {
			entity = entities.IntelligenceEnum(code).ToString()
		}
	case "control":
		if code > 4 {
			return ""
		} else {
			entity = entities.ControlClassificationEnum(code).ToString()
		}
	case "movement":
		if code > 8 {
			return ""
		} else {
			entity = entities.MovementEnum(code).ToString()
		}
	case "basic_attack":
		if code > 9 {
			return ""
		} else {
			entities.BasicAttackEnum(code).ToString()
		}
	case "ability":
		if code > 6{
			return ""
		} else {
			entities.AbilityEnum(code).ToString()
		}
	case "major_category":
		if code > 3 {
			return ""
		} else {
			entity = entities.MajorCategoryEnum(code).ToString()
		}
	case "dmg_type":
		if code > 4 {
			return ""
		} else {
			entity = entities.DmgTypeEnum(code).ToString()
		}
	case "class":
		if code > 3 {
			return ""
		} else {
			entity = entities.ClassEnum(code).ToString()
		}
	case "gender":
		if code > 3 {
			return ""
		} else {
			entity = entities.GenderEnum(code).ToString()
		}
	default:
		return ""
	}
	return entity
}