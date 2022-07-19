package middlewares

import (
	"net/http"

	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/gin-gonic/gin"
)

// The argument must first be initialized at "setup/routes/routes.go" to initialize its struct values
func AuthorizeJWT(j services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// fetch the auth header received from frontend requests
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			responseJSON := responses.ErrorResponse("Failed to process jwt request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, responseJSON)
			return
		}
		// This is merely a function caller
		// validate it to make sure that the user has signed-in before requesting the contents it is protecting
		token := j.ValidateToken(authHeader, c)
		if !token.Valid || token == nil {
			response := responses.ErrorResponse("JWT error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}