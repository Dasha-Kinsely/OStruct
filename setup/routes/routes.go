package routes

import (
	//"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	//"github.com/dasha-kinsely/ostruct/setup/config"

	"github.com/dasha-kinsely/ostruct/controllers/handlers"
	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/dasha-kinsely/ostruct/setup/config"
	"github.com/gin-gonic/gin"
)

// The following var are global and should be useful throughout the current user session
var (
	db = config.GetSqlDB()
	userRepo repos.UserRepo = repos.NewUserRepo(db)
	jwtService services.JWTService = services.NewJWTService()
	authService services.AuthService = services.NewAuthService(userRepo)
	userService services.UserService = services.NewUserService(userRepo)
	authHandler handlers.AuthHandler = handlers.NewAuthHandler(authService, jwtService, userService)
)

func InitRoutes(r *gin.Engine) {
	// routes that do not require middlewares
	base := r.Group("/api")
	{
		base.GET("/signup", authHandler.Signup)
		base.GET("/signin", func(c *gin.Context){
			return
		})
		base.GET("/signout", func(c *gin.Context){
			return
		})
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
