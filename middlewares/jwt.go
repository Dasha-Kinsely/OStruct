package middlewares

import (
	"net/http"

	"github.com/dasha-kinsely/ostruct/controllers/helpers"
	"github.com/dasha-kinsely/ostruct/controllers/responses"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(token *helpers.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			responseJSON := responses.ErrorResponse("Failed to process jwt request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, responseJSON)
			return
		}
		token := helpers.JWTService.ValidateToken(authHeader, c)
		//token := helpers.JWTService.ValidateToken(authHeader, c)
		if !token.Valid {
			response := responses.ErrorResponse("JWT error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}