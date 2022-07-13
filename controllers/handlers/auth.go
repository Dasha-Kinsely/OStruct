package handlers

import "github.com/gin-gonic/gin"

type Auth interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}
