package logic

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hash string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ComparePassword(hashPassword string, input string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(input))
	return err
}

func GenToken(secretKey string, iat, seconds int64, userID int64) (tokenGen string, err error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userID"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
