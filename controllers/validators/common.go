package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Bind(c *gin.Context, obj interface{}) error {
	return c.ShouldBindWith(obj, binding.Default(c.Request.Method, c.ContentType()))
}