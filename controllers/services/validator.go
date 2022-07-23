package services

import (
	"net/http"

	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/models/commonstructs"
	"github.com/gin-gonic/gin"
)

type ValidatorService interface {
	Validate(c *gin.Context, err error, step string) bool
} 
type validatorService struct {
}

func NewValidatorService() ValidatorService {
	return &validatorService{}
}

func (v validatorService) Validate(c *gin.Context, err error, step string) bool {
	empty := commonstructs.Empty{}
	switch step {
	case "signupForm":
		if err != nil {
			res := responses.ErrorResponse("sign up form is in an invalid format", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return false
		} else {
			return true
		}
	case "createUser":
		if err != nil {
			res := responses.ErrorResponse("problem occurred while signing up user", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	case "signinForm":
		if err != nil {
			res := responses.ErrorResponse("sign in form is in an invalid format", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return false
		} else {
			return true
		}
	case "signinCred" :
		if err != nil {
			res := responses.ErrorResponse("problem occurred while signing in user", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
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