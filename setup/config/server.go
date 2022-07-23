package config

import (
	"github.com/dasha-kinsely/ostruct/controllers/handlers"
	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeServer(server *gin.Engine) {
	LoadEnv()
	// The following var are global and should be useful throughout the current user session
	var (
		db *gorm.DB = SetupSqlDBConnection()
		userRepo repos.UserRepo = repos.NewUserRepo(db)
		jwtService services.JWTService = services.NewJWTService()
		authService services.AuthService = services.NewAuthService(userRepo)
		userService services.UserService = services.NewUserService(userRepo)
		validatorService services.ValidatorService = services.NewValidatorService()
		authHandler handlers.AuthHandler = handlers.NewAuthHandler(authService, jwtService, userService, validatorService)
	)
	if reqActions := StartupCheck(); reqActions == true {
		BulkMigrate(db)
	}
	SetupMongoDBConnection()
	SetupRedisConnection()
	// routes that do not require middlewares
	base := server.Group("/api")
	{
		base.GET("/signup", authHandler.Signup)
		base.GET("/signin", authHandler.Signin)
		base.GET("/signout", func(c *gin.Context){
			return
		})
	}
	// routes that require middlewares
	safe := server.Group("/authorized", middlewares.AuthorizeJWT(jwtService))
	{
		safe.GET("/", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "all good Kim",
			})
		})
	}
}