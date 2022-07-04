package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"todo-list-api/config"
)

func Get(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Unix() + 1200,
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(config.JWT.Key)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func Parse(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return config.JWT.Key, nil
	})

	return t, err
}

func ParseClaims(token *jwt.Token) map[string]interface{} {
	return token.Claims.(jwt.MapClaims)
}
