package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type JWTClaim struct {
	AdminID uint `json:"admin_id"`
	jwt.StandardClaims
}

func GenerateJWT(adminID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		AdminID: adminID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
