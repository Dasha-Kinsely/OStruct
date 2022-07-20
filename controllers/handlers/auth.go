package handlers

import (
	"net/http"

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
	authService services.AuthService
	jwtService services.JWTService
	userService services.UserService
}

func NewAuthHandler(
	authService services.AuthService,
	jwtService services.JWTService,
	userService services.UserService,
	) AuthHandler {
		return &authHandler{
			authService: authService,
			jwtService: jwtService,
			userService: userService,
		}
}

func (handler *authHandler) Signup(c *gin.Context) {
	// check if the incoming form request is in valid format
	var signupRequest dto.SignupRequest
	if err := c.ShouldBind(&signupRequest); err != nil {
		res := responses.ErrorResponse("sign up form is in an invalid format", err.Error(), others.Empty{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return 
	}
	// create the user by accessing userService interface
	user, err := handler.userService.CreateUser(signupRequest)
	if err != nil {
		res := responses.ErrorResponse("problem occurred while signing up user", err.Error(), others.Empty{})
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	responses.NewSignupSerializer(c, user)
}

func (handler *authHandler) Signin(c *gin.Context) {
	// var signinRequest dto.SigninRequest
	return
}

func (handler *authHandler) Signout(c *gin.Context) {
	return
}