package setup

import (
	"os"

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
	// does this run require setting up database schemas
	if os.Getenv("REQUIRES_MIGRATION") == "yes" {
		BulkMigrate()
	}
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
}

func BulkMigrate() {
	//migration.MigrateMongoDB()
	migration.MigrateSqlDB()
}