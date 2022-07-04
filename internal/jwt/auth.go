package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"todo-list-api/config"
	"todo-list-api/pkg/logger"
)

func Auth(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie(config.JWT.Cookie)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, fmt.Sprintf("no %s cookie", config.JWT.Cookie), 401)
		return nil, err
	}

	token := cookie.Value

	// TODO switch on error and make certain messages
	t, err := Parse(token)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "jwt is invalid", 401)
		return nil, err
	}

	return t, nil
}
