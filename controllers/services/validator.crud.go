package services

import (
	"net/http"

	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/models/commonstructs"
	"github.com/gin-gonic/gin"
)

type ValidatorCRUDService interface {
	Validate(c *gin.Context, err error, step string) bool
}

type validatorCRUDService struct {
}

func NewCRUDValidatorService() ValidatorCRUDService {
	return &validatorCRUDService{}
}

func (v validatorCRUDService) Validate(c *gin.Context, err error, step string) bool {
	empty := commonstructs.Empty{}
	switch step{
	case "form":
		if err != nil {
			res := responses.ErrorResponse("your request form is not in a valid format", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return false
		} else {
			return true
		}
	case "insert":
		if err != nil {
			res := responses.ErrorResponse("problem with object insertion", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	case "fetch":
		if err != nil {
			res := responses.ErrorResponse("problem with object fetching", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	case "edit":
		if err != nil {
			res := responses.ErrorResponse("problem with object editing", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	case "delete":
		if err != nil {
			res := responses.ErrorResponse("problem with object deletion", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	case "enum":
		res := responses.ErrorResponse("enum is invalid", err.Error(), empty)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return false
	default:
		if err != nil {
			res := responses.ErrorResponse("this is an unknown error, please see error body for details", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	}
}