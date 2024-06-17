package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ecdsaPrivateKey *ecdsa.PrivateKey

// Инициализация ключа
func init() {
	var err error
	ecdsaPrivateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("error generating ECDSA key: %v", err)
	}
}

func CreateToken(id int, username, role string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
