package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CRUDSerializer struct {
	Message string `json:"msg"`
	Payload interface{} `json:"payload"`
}

func NewCRUDSerializer(c *gin.Context, payload interface{}, msg string) {
	serializer := CRUDSerializer{
		Message: msg,
		Payload: payload,
	}
	c.JSON(http.StatusOK, serializer)
	return
}