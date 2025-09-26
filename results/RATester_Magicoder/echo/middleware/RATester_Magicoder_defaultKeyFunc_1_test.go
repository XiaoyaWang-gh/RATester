package middleware

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestDefaultKeyFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("my-secret-key"),
	}

	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(config.SigningKey)

	parsedToken, err := jwt.Parse(tokenString, config.defaultKeyFunc)
	if err != nil {
		t.Errorf("Error parsing token: %v", err)
	}

	if !parsedToken.Valid {
		t.Error("Parsed token is not valid")
	}
}
