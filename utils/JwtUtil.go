package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Key = []byte("itheima")
)

// GenToken generates a JWT token with the provided claims
func GenToken(claims jwt.MapClaims) (string, error) {
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tString, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}
	return tString, nil
}

// ParseToken parses a JWT token and returns the claims
func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["claims"].(map[string]interface{}), nil
	} else {
		return nil, fmt.Errorf("invalid claims type")
	}
}
