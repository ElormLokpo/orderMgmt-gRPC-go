package bcrypt

import (
	"auth-service/errHandler"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	errHandler.StdErr(err)

	return string(hashedPassword)
}

func CompareHashedPasswords(password string, hasehdPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasehdPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
