package jwt

import (
	"auth-service/errHandler"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	Userid string
}

func GenerateToken(userid string) string {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Elorm Lokpo",
			ExpiresAt: time.Now().Add(time.Hour * 500).Unix(),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tkn.SignedString("SHAMBALA")
	errHandler.StdErr(err)

	return token
}
