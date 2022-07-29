package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindParamInRequest(c *gin.Context, param string) string {
	parameter := c.Param(param)
	fmt.Println("Param is: "+parameter)
	if parameter == "" {
		msg := fmt.Sprintf("cannot find \"%s\" in your url parameters", param)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": msg,
		}) 
		return ""
	} else {
		return parameter
	}
}