package config

import (
	"github.com/dasha-kinsely/ostruct/controllers/handlers"
	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/controllers/services"
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func InitializeServer(server *gin.Engine) {
	LoadEnv()
	// The following var are global and should be useful throughout the current user session
	var (
		// databases
		db *gorm.DB = SetupSqlDBConnection()
		fdb *mongo.Client = SetupMongoDBConnection()
		cache *redis.Client = SetupRedisConnection()
		// repo access layer, requires at least one databases as argument
		userRepo repos.UserRepo = repos.NewUserRepo(db)
		equipmentRepo repos.EquipmentRepo = repos.NewEquipmentRepo(db)
		neutralRepo repos.NeutralRepo = repos.NewNeutralRepo(db)
		// services layer, cannot access db directly
		jwtService services.JWTService = services.NewJWTService()
		authService services.AuthService = services.NewAuthService(userRepo)
		userService services.UserService = services.NewUserService(userRepo)
		equipmentService services.EquipmentCRUDService = services.NewEquipmentService(equipmentRepo)
		neutralService services.NeutralCRUDService = services.NewNeutralCRUDService(neutralRepo)
		validatorAuthService services.ValidatorAuthService = services.NewAuthValidatorService()
		validatorCRUDService services.ValidatorCRUDService = services.NewCRUDValidatorService()
		// routers' handler layer, cannot access db directly
		authHandler handlers.AuthHandler = handlers.NewAuthHandler(authService, jwtService, userService, validatorAuthService)
		neutralCRUDHandler handlers.NeutralCRUDHandler = handlers.NewNeutralCRUD(jwtService, neutralService, validatorCRUDService)
		equipmentCRUDHandler handlers.EquipmentHandler = handlers.NewEquipmentHandler(equipmentService, jwtService, validatorCRUDService)
	)
	if StartupCheck() {
		BulkMigrate(db, fdb, cache)
	}
	// routes that do not require middlewares
	base := server.Group("/api")
	{
		base.GET("/signup", authHandler.Signup)
		base.GET("/signin", authHandler.Signin)
		//--------------------------------------------------------------------------------------------
		spellschool := base.Group("/spell-school") 
		{
			spellschool.POST("/create")
			spellschool.GET("/get-all-omit-abilities")
			spellschool.GET("/get-details/:key")
			spellschool.PUT("modify-base")
			spellschool.PUT("/modify-details/add")
			spellschool.PUT("/modify-details/edit")
			spellschool.PUT("/modify-details/remove")
			spellschool.DELETE("/del/:name")
			spellschool.GET("statistic/view") // this is updated by a smart-trigger from mongo_db related routes
		}
		techtype := base.Group("/tech-type")
		{
			techtype.POST("/create")
			techtype.GET("/get-all-omit-abilities")
			techtype.GET("/get-details/:key")
			techtype.PUT("/modify-base")
			techtype.PUT("/modify-details/add")
			techtype.PUT("/modify-details/edit")
			techtype.PUT("/modify-details/remove")
			techtype.DELETE("/del/:name")
			techtype.GET("statistic/view") // this is updated by a smart-trigger from mongo_db related routes
		}
		//--------------------------------------------------------------------------------------------
		cc := base.Group("/cc")
		{
			cc.POST("/create")
			cc.PUT("/edit/:name")
			cc.GET("/get-table")
			cc.GET("/get/:name")
			cc.GET("/filter/default") // dto_smart_array
			cc.GET("/filter/mitigation") // dto
			cc.DELETE("/del/:name")
			cc.GET("statistic/view") // this is updated by a smart-trigger from mongo_db related routes
		}
		//--------------------------------------------------------------------------------------------
		equipments := base.Group("/equipments")
		{
			equipments.POST("/create", equipmentCRUDHandler.CreateEquipment)
			equipments.GET("/get-all", equipmentCRUDHandler.GetAllEquipment)
			equipments.GET("/get-by-category/:category", equipmentCRUDHandler.GetAllEquipmentByCategory)
			equipments.GET("/get/:name", equipmentCRUDHandler.GetOneEquipment)
			equipments.PUT("/edit/:name", equipmentCRUDHandler.UpdateOneEquipment)
			equipments.PUT("/definition-update", equipmentCRUDHandler.UpdateEquipLimit)
			equipments.DELETE("/del/:name", equipmentCRUDHandler.DeleteOneEquipment)
		}
		//--------------------------------------------------------------------------------------------
		obj := base.Group("/obj")
		{
			obj.POST("/neutral-monsters/create", neutralCRUDHandler.CreateNeutralMonster)
			obj.GET("/neutral-monsters/get-all", neutralCRUDHandler.GetAllNeutralMonster)
			obj.GET("/neutral-monsters/get/:name", neutralCRUDHandler.GetOneNeutralMonster)
			obj.PUT("/neutral-monsters/edit", neutralCRUDHandler.UpdateOneNeutralMonster)
			obj.DELETE("/neutral-monsters/del/:name", neutralCRUDHandler.DeleteOneNeutralMonster)
			//--------------------------------------------------------------------------------------------
			obj.POST("/subordinate-monsters/create", neutralCRUDHandler.CreateSubordinateMonster)
			obj.GET("/subordinate-monsters/get-all", neutralCRUDHandler.GetAllSubordinateMonster)
			obj.GET("/subordinate-monsters/get/:name", neutralCRUDHandler.GetOneSubordinateMonster)
			obj.PUT("/subordinate-monsters/edit/:name", neutralCRUDHandler.UpdateOneSubordinateMonster)
			obj.DELETE("/subordinate-monsters/del/:name", neutralCRUDHandler.DeleteOneSubordinateMonster)
		}
		//--------------------------------------------------------------------------------------------
	}
	// routes that require middlewares
	safe := server.Group("/authorized", middlewares.AuthorizeJWT(jwtService))
	{
		// TODO:
		safe.GET("/profile", func(c *gin.Context){
			return
		})
		safe.GET("/signout", func(c *gin.Context){
			return
		})
	}
}

/*
		base.POST("/create")
		base.GET("/get-all")
		base.GET("/get/:name")
		base.PUT("/edit/:name")
		base.DELETE("/del/:name")
*/