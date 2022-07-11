package config

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)


func LoadEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalf("failed to find .env files !!!, %v", err)
	} else {
		fmt.Println("env variables loaded successfully...")
	}
}