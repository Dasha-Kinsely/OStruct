package setup

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *gin.Engine
}

func LoadEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalf("failed to find .env files !!!, %v", err)
	} else {
		fmt.Println("env variables loaded successfully...")
	}
}

func Run() {
	LoadEnv()
	fmt.Fprintln(os.Getenv("DB"))
}