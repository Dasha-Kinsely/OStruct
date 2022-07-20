package responses

import (
	"net/http"

	"github.com/dasha-kinsely/ostruct/models/entities"
	"github.com/gin-gonic/gin"
)

type SignupSerializer struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func NewSignupSerializer(c *gin.Context, user entities.User) {
	serializer := SignupSerializer{
		ID: int64(user.ID),
		Username: user.Username,
		Email: user.Email,
	}
	c.JSON(http.StatusCreated, serializer)
	return
}