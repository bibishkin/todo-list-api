package task

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
	"todo-list-api/config"
	"todo-list-api/internal/jwt"
	"todo-list-api/pkg/logger"
)

func Auth(w http.ResponseWriter, r *http.Request) (*Task, error) {
	token, err := jwt.Auth(w, r)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "400 bad task id", 400)
		return nil, err
	}

	task, err := GetTaskByID(context.Background(), config.DBPool, id)
	if err != nil {
		logger.InfoLog.Println(err)
		if err == pgx.ErrNoRows {
			http.Error(w, "404 task not found", 404)
			return nil, err
		}
		http.Error(w, "500 internal server error", 500)
		return nil, err
	}

	m := jwt.ParseClaims(token)
	userID := int(m["sub"].(float64))
	if task.UserID != userID {
		err = errors.New("forbidden")
		logger.InfoLog.Println(err)
		http.Error(w, "403 forbidden", 403)
		return nil, err
	}

	return task, nil
}
