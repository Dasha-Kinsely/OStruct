package handlers

import (
	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/utils"
	"github.com/gin-gonic/gin"
)

type NeutralCRUDHandler interface {
	CreateNeutralMonster(c *gin.Context)
	GetAllNeutralMonster(c *gin.Context)
	GetOneNeutralMonster(c *gin.Context)
	UpdateOneNeutralMonster(c *gin.Context)
	DeleteOneNeutralMonster(c *gin.Context)
	//------------------------------------------
	CreateSubordinateMonster(c *gin.Context)
	GetAllSubordinateMonster(c *gin.Context)
	GetOneSubordinateMonster(c *gin.Context)
	UpdateOneSubordinateMonster(c *gin.Context)
	DeleteOneSubordinateMonster(c *gin.Context)
}

type neutralCRUDHandler struct {
	jwtService services.JWTService
	neutralService services.NeutralCRUDService
	validatorCRUD services.ValidatorCRUDService
}

func NewNeutralCRUD(jwt services.JWTService, neutralCRUDService services.NeutralCRUDService, validator services.ValidatorCRUDService) NeutralCRUDHandler {
	return &neutralCRUDHandler{
		jwtService: jwt,
		neutralService: neutralCRUDService,
		validatorCRUD: validator,
	}
}

func (handler *neutralCRUDHandler) CreateNeutralMonster(c *gin.Context) {
	var req dto.CreateNeutralMonsterRequest
	err := c.ShouldBind(&req)
	if handler.validatorCRUD.Validate(c, err, "form") {
		o, err := handler.neutralService.CreateNeutralMonster(req)
		if handler.validatorCRUD.Validate(c, err, "insert") {
			responses.NewCRUDSerializer(c, o, "inserted the following neutral monster")
		}
	}
}

func (handler *neutralCRUDHandler) GetAllNeutralMonster(c *gin.Context) {
	oList, err := handler.neutralService.GetAllNeutralMonster()
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, oList, "found the follwoing neutral monsters")
	}
}

func (handler *neutralCRUDHandler) GetOneNeutralMonster(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == ""{
		return
	}
	o, err := handler.neutralService.GetOneNeutralMonster(name)
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, o, "found this neutral monster")
	}
}

func (handler *neutralCRUDHandler) UpdateOneNeutralMonster(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == ""{
		return
	}
	var req dto.UpdateOneNeutralMonster
	err := c.ShouldBind(&req)
	if handler.validatorCRUD.Validate(c, err, "form") {
		o, err := handler.neutralService.UpdateOneNeutralMonster(req, name) 
		if handler.validatorCRUD.Validate(c, err, "edit") {
			responses.NewCRUDSerializer(c, o, "edited the following neutral monster")
		}
	}
}

func (handler *neutralCRUDHandler) DeleteOneNeutralMonster(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == "" {
		return
	}
	err := handler.neutralService.DeleteOneNeutralMonster(name)
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, nil, "deletion successful")
	}
}

// ------------------------------------------------------------------------------------
func (handler *neutralCRUDHandler) CreateSubordinateMonster(c *gin.Context) {
	var req dto.CreateSubordinateMonsterRequest
	err := c.ShouldBind(&req)
	if handler.validatorCRUD.Validate(c, err, "form") {
		o, err := handler.neutralService.CreateSubordinateMonster(req)
		if handler.validatorCRUD.Validate(c, err, "insert"){
			responses.NewCRUDSerializer(c, o, "inserted the following object")
		}
	}
}

func (handler *neutralCRUDHandler) GetAllSubordinateMonster(c *gin.Context) {
	oList, err := handler.neutralService.GetAllSubordinateMonster()
	if handler.validatorCRUD.Validate(c, err, "fetch"){
		responses.NewCRUDSerializer(c, oList, "found the following objects")
	}
}

func (handler *neutralCRUDHandler) GetOneSubordinateMonster(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == "" {
		return
	}
	o, err := handler.neutralService.GetOneSubordinateMonster(name)
	if handler.validatorCRUD.Validate(c, err, "fetch") {
		responses.NewCRUDSerializer(c, o, "found the following object")
	}
}

func (handler *neutralCRUDHandler) UpdateOneSubordinateMonster(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == ""{
		return
	}
	var req dto.UpdateOneSubordinateMonster
	err := c.ShouldBind(&req)
	if handler.validatorCRUD.Validate(c, err, "form") {
		o, err := handler.neutralService.UpdateOneSubordinateMonster(req, name) 
		if handler.validatorCRUD.Validate(c, err, "edit") {
			responses.NewCRUDSerializer(c, o, "edited the following object")
		}
	}
}

func (handler *neutralCRUDHandler) DeleteOneSubordinateMonster(c *gin.Context) {
	name := utils.FindParamInRequest(c, "name")
	if name == "" {
		return
	}
	err := handler.neutralService.DeleteOneSubordinateMonster(name)
	if handler.validatorCRUD.Validate(c, err, "delete") {
		responses.NewCRUDSerializer(c, nil, "deletion successful")
	}
}