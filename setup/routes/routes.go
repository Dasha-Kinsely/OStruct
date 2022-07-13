package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// routes that do not require middlewares
	base := r.Group("/api")
	{
		base.GET("/test", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "all good Saul",
			})
		})
	}
	// routes that require middlewares
	
}