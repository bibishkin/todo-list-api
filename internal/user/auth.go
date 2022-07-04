package user

import (
	"context"
	"github.com/jackc/pgx/v4"
	"io"
	"net/http"
	"todo-list-api/config"
	"todo-list-api/internal/jwt"
	"todo-list-api/pkg/logger"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	u, err := parseUser(r)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "bad request", 400)
		return
	}

	u2, err := GetUserByEmail(context.Background(), config.DBPool, u.Email)
	if err == pgx.ErrNoRows {
		logger.InfoLog.Println(err)
		http.Error(w, "user doesn't exist", 404)
		return
	}

	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	if u2.Password != u.Password {
		logger.InfoLog.Println(err)
		http.Error(w, "wrong password", 400)
		return
	}

	token, err := jwt.Get(u2.ID)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: token})
	w.WriteHeader(200)
	return
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	u, err := parseUser(r)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "bad request", 400)
		return
	}

	_, err = GetUserByEmail(context.Background(), config.DBPool, u.Email)

	if err != pgx.ErrNoRows {
		if err == nil {
			logger.InfoLog.Println(err)
			http.Error(w, "user already exist", 409)
			return
		}
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	err = InsertUser(context.Background(), config.DBPool, u)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	w.WriteHeader(201)
	return
}

func parseUser(r *http.Request) (*User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	u, err := Unmarshal(body)
	if err != nil {
		return nil, err
	}
	return u, nil
}
