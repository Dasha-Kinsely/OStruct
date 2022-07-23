package setup

import (
	"github.com/dasha-kinsely/ostruct/middlewares"
	"github.com/dasha-kinsely/ostruct/setup/config"
	"github.com/gin-gonic/gin"
)

func Run() {
	// get the server ready
	server := gin.Default()
	config.InitializeServer(server)
	server.Use(middlewares.CorsMiddleware())
	server.Run()
}

