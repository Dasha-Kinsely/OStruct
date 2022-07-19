package routes

import (
	//"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	//"github.com/dasha-kinsely/ostruct/setup/config"

	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/gin-gonic/gin"
)

// The following var are global and should be useful throughout the current user session
var (
	jwtService services.JWTService = services.NewJWTService()
)

func InitRoutes(r *gin.Engine) {
	// routes that do not require middlewares
	base := r.Group("/api")
	{
		base.GET("/signin", func(c *gin.Context){
			return
		})
		base.GET("/signup")
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
/*
base.GET("/test", func(c *gin.Context){
			clientR := config.GetRedis()
			pong, err := clientR.Ping(c).Result()
			fmt.Println(pong, err)
			clientM, _ := config.GetMongoClient().ListDatabaseNames(c, bson.D{{}})
			fmt.Printf("mongo db names are %s\n", clientM)
			clientS, _ := config.GetSqlDB().DB()
			fmt.Printf("mysql db names are %s\n", clientS)
			c.JSON(200, gin.H{
				"message": "all good Saul",
			})
		})
*/