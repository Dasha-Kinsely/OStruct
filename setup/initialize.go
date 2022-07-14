package setup

import (
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/dasha-kinsely/ostruct/setup/config"
	"github.com/dasha-kinsely/ostruct/setup/migration"
	"github.com/dasha-kinsely/ostruct/setup/routes"
	"github.com/gin-gonic/gin"
)

func OnStartup() {
	config.LoadEnv()
	config.SetupSqlDBConnection()
	config.SetupMongoDBConnection()
	config.SetupRedisConnection()
}

func Run() {
	// set up of databases & port
	OnStartup()
	defer config.CloseSqlDBConnection(config.SqlDB)
	// get the server ready
	server := gin.Default()
	server.Use(middlewares.CorsMiddleware())
	routes.InitRoutes(server)
	server.Run()
	migration.MigrateMongoDB()
}