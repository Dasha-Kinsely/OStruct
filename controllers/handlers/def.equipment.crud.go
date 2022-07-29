package handlers

import (
	"strconv"

	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/utils"
	"github.com/gin-gonic/gin"
)

type EquipmentHandler interface {
	CreateEquipment(c *gin.Context)
	GetAllEquipment(c *gin.Context)
	GetAllEquipmentByCategory(c *gin.Context)
	GetOneEquipment(c *gin.Context)
	UpdateOneEquipment(c *gin.Context)
	UpdateEquipLimit(c *gin.Context)
	DeleteOneEquipment(c *gin.Context)
}

type equipmentHandler struct {
	equipmentService services.EquipmentCRUDService
	jwt services.JWTService
	validatorCRUD services.ValidatorCRUDService
}

func NewEquipmentHandler(crud services.EquipmentCRUDService, jwt services.JWTService, validatorCRUD services.ValidatorCRUDService) EquipmentHandler {
	return &equipmentHandler{
		equipmentService: crud,
		jwt: jwt,
		validatorCRUD: validatorCRUD,
	}
}

func (handler *equipmentHandler) CreateEquipment(c *gin.Context) {
	var req dto.CreateEquipmentRequest
	err := c.ShouldBind(&req)
	if handler.validatorCRUD.Validate(c, err, "form") {
		parsed, err := strconv.ParseInt(req.Category, 10, 8)
		if err != nil {
			responses.ErrorResponse("failed to interpret enum", "Hacker Warning: the handler has failed due to errors that originated from client side !!!", nil)
			return
		}
		if interpreted := utils.InterpretEnum("equipment", int8(parsed)); interpreted == "" {
			handler.validatorCRUD.Validate(c, nil, "enum")
		}
		o, err := handler.equipmentService.CreateEquipment(req)
		if handler.validatorCRUD.Validate(c, err, "insert") {
			responses.NewCRUDSerializer(c, o, "inserted the follwoing equipment definition")
		}
	}
}

func (handler *equipmentHandler) GetAllEquipment(c *gin.Context) {
	oList, err := handler.equipmentService.GetAllEquipment()
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, oList, "found the following equipments")
	}
}

func (handler *equipmentHandler) GetAllEquipmentByCategory(c *gin.Context) {
	category := utils.FindParamInRequest(c, "category")
	if category == ""{
		return
	}
	oList, err := handler.equipmentService.GetAllEquipmentByCategory(category)
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, oList, "found the following equipments")
	}
}

func (handler *equipmentHandler) GetOneEquipment(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == ""{
		return
	}
	o, err := handler.equipmentService.GetOneEquipment(name)
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, o, "found the following equipment")
	}
}

func (handler *equipmentHandler) UpdateOneEquipment(c *gin.Context) {
	var req dto.UpdateOneEquipmentRequest
	name := utils.FindParamInRequest(c, "name")
	if name == ""{
		return
	}
	o, err := handler.equipmentService.UpdateOneEquipment(req, name)
	if handler.validatorCRUD.Validate(c, err, "edit") {
		responses.NewCRUDSerializer(c, o, "updated the following object")
	}
}

func (handler *equipmentHandler) UpdateEquipLimit(c *gin.Context) {
	var req dto.BulkUpdateEquipLimitRequest
	err := handler.equipmentService.UpdateEquipLimit(req)
	if handler.validatorCRUD.Validate(c, err, "edit") {
		responses.NewCRUDSerializer(c, nil, "updated equipment limit")
	}
}

func (handler *equipmentHandler) DeleteOneEquipment(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == ""{
		return
	}
	err := handler.equipmentService.DeleteOneEquipment(name)
	if handler.validatorCRUD.Validate(c, err, "delete") {
		responses.NewCRUDSerializer(c, nil, "deletion successful")
	}
}