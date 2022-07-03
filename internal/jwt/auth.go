package jwt

import (
	"fmt"
	"net/http"
	"todo-list-api/config"
	"todo-list-api/pkg/logger"
)

func Auth(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(config.JWT.Cookie)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, fmt.Sprintf("no %s cookie", config.JWT.Cookie), 401)
		return err
	}

	token := cookie.Value

	// TODO switch on error and make certain messages
	_, err = Parse(token)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "jwt is invalid", 401)
		return err
	}

	return nil
}
