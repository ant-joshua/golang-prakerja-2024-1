package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	ID       int    `json:"id"`
}

func GenerateJwtToken(username, email string, id int) (string, error) {
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": JwtCustomClaims{
			Email:    email,
			Username: username,
			ID:       id,
		},
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Printf("Error when generate jwt token: %v", err)

		return "", err
	}

	return tokenString, nil
}

func VerifyJwtToken(tokenString string) (*JwtCustomClaims, error) {
	secret := "secret"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Printf("Error when verify jwt token: %v", err)

		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		log.Printf("Error when verify jwt token: %v", err)

		return nil, err
	}

	data := claims["data"].(map[string]interface{})

	return &JwtCustomClaims{
		Email:    data["email"].(string),
		Username: data["username"].(string),
		ID:       int(data["id"].(float64)),
	}, nil
}
