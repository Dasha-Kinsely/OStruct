package handlers

import (
	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/models/others"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
	Signout(c *gin.Context)
}
// bound to its corresponding interface
type authHandler struct {
	auth services.AuthService
	jwt services.JWTService
	user services.UserService
}

func NewAuthHandler(
	auth services.AuthService,
	jwt services.JWTService,
	user services.UserService,
	) AuthHandler {
		return &authHandler{
			auth: auth,
			jwt: jwt,
			user: user,
		}
}

func (handler *authHandler) Signup(c *gin.Context) {
	var signupRequest dto.SignupRequest
	if err := c.ShouldBind(&signupRequest); err != nil {
		res := responses.ErrorResponse("sign up form is in an invalid format", err.Error(), others.Empty)
		return 
	}

}

func (handler *authHandler) Signin(c *gin.Context) {
	var signinRequest dto.SigninRequest

}

func (handler *authHandler) Signout(c *gin.Context) {
	
}