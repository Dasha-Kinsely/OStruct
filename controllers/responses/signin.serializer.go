package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SigninSerializer struct {
	Email string `json:"email"`
}

func NewSigninSerializer(c *gin.Context, email string) {
	serializer := SigninSerializer{
		Email: email,
	}
	c.JSON(http.StatusCreated, serializer)
	return
}