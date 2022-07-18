package routes

import (
	"fmt"

	"github.com/dasha-kinsely/ostruct/controllers/helpers"
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/dasha-kinsely/ostruct/setup/config"
	"github.com/gin-gonic/gin"
)

var (
	jwtService helpers.JWTService = helpers.NewJWTService()
)

func InitRoutes(r *gin.Engine) {
	// routes that do not require middlewares
	base := r.Group("/api")
	{
		base.GET("/test", func(c *gin.Context){
			client := config.GetRedis()
			pong, err := client.Ping(c).Result()
			fmt.Println(pong, err)
			c.JSON(200, gin.H{
				"message": "all good Saul",
			})
		})
		//base.GET("/public", )
	}
	// routes that require middlewares
	safe := r.Group("/authorized", middlewares.AuthorizeJWT(jwtService))
	{
		safe.GET("/", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "all good Kim",
			})
		})
	}
}