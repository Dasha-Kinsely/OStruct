package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) string {
	bytePassword := []byte(pass)
	hashed, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashed)
}

func DecryptPassword(formPassword string, actualPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(formPassword))
}