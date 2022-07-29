package handlers

import (
	"log"
	"strconv"

	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/models/dto"
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
	validatorAuthService services.ValidatorAuthService
}

func NewAuthHandler(
	authService services.AuthService,
	jwtService services.JWTService,
	userService services.UserService,
	validatorAuthService services.ValidatorAuthService,
	) AuthHandler {
		return &authHandler{
			authService: authService,
			jwtService: jwtService,
			userService: userService,
			validatorAuthService: validatorAuthService,
		}
}

func (handler *authHandler) Signup(c *gin.Context) {
	// check if the incoming form request is in valid format
	var signupRequest dto.SignupRequest
	err := c.ShouldBind(&signupRequest)
	if handler.validatorAuthService.Validate(c, err, "signupForm") {
		// create the user by accessing userService interface
		user, err := handler.userService.CreateUser(signupRequest)
		if handler.validatorAuthService.Validate(c, err, "createUser") {
			responses.NewSignupSerializer(c, user)
		}
	}
	return
}

func (handler *authHandler) Signin(c *gin.Context) {
	// check if the incoming form request is in valid format
	var signinRequest dto.SigninRequest
	err := c.ShouldBind(&signinRequest)
	if handler.validatorAuthService.Validate(c, err, "signinForm") {
		// verify user credentials
		userID, err := handler.authService.VerifyCredentials(signinRequest.Email, signinRequest.Password) 
		if handler.validatorAuthService.Validate(c, err, "signinCred") {
			token := handler.jwtService.GenerateToken(strconv.FormatUint(uint64(userID), 10))
			c.Set("authenticated", token)
			c.Set("user", signinRequest.Email)
			log.Println(c.Get("authenticated"))
			responses.NewSigninSerializer(c, signinRequest.Email)
		}
	}
	return
}

func (handler *authHandler) Signout(c *gin.Context) {
	return
}