package jwt

import (
	"fmt"
	"log"
	"time"

	errorhandlers "github.com/Samito19/msu-mp-auth-ms/errorHandlers"
	"github.com/golang-jwt/jwt/v4"
)

func CreateNewToken() (string, error) {
	defer func() {
		if tokenErr := recover(); tokenErr != nil {
			log.Printf("Runtime panic: %v", tokenErr)
		}
	}()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hello": "world",
		"exp":   time.Now().UTC().Unix(),
	})

	tokenString, tokenErr := token.SignedString([]byte("secret"))

	return tokenString, tokenErr
}

func VerifyToken(tokenString string) {
	defer func() {
		if parseErr := recover(); parseErr != nil {
			log.Printf("Runtime panic: %v", parseErr)
		}
	}()

	token, parseErr := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	errorhandlers.CheckError(parseErr)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("hello : %s", claims["hello"])
	} else {
		fmt.Println("Invalid token")
	}
}
