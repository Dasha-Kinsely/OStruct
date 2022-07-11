package setup

import (
	"github.com/dasha-kinsely/ostruct/setup/config"
	"github.com/dasha-kinsely/ostruct/setup/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	// set up of databases & port
	config.LoadEnv()
	config.SetupSqlDBConnection()
	defer config.CloseSqlDBConnection(config.DB)
	// get the server ready
	server := gin.Default()
	routes.InitRoutes(server)
	server.Run()
}