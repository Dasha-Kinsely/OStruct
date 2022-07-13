package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	access_whitelist := os.Getenv("CORS_WHITELIST")
	access_methods := os.Getenv("CORS_METHODS")
	access_headers := os.Getenv("CORS_HEADERS")
	return func(c *gin.Context){
		// setup request restrctions
		c.Writer.Header().Set("Access-Control-Allow-Origin", access_whitelist)
		c.Writer.Header().Set("Access-Control-Allow-Methods", access_methods)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", access_headers)
		// skip options method
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}