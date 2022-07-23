package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) string {
	bytePassword := []byte(pass)
	fmt.Println(bytePassword)
	fmt.Println("Encryption: "+pass)
	hashed, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashed)
}

func DecryptPassword(hashedPassword string, formPassword string) error {
	hashed := []byte(hashedPassword)
	form := []byte(formPassword)
	return bcrypt.CompareHashAndPassword(hashed, form)
}