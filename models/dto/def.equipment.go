package dto

type CreateEquipmentRequest struct {
	EquipmentType string `json:"equipmenttype" form:"equipmenttype" binding:"required"`
	Category string `json:"category" form:"category" binding:"required"`
}

type UpdateOneEquipmentRequest struct {
	EquipmentType string `json:"equipmenttype" form:"equipmenttype" binding:"required"`
	Category string `json:"category" form:"category" binding:"required"`
}

type BulkUpdateEquipLimitRequest struct {
	Category string `json:"category" form:"category" binding:"required"`
	EquipLimit string `json:"equiplimit" form:"equiplimit" binding:"required"`
}