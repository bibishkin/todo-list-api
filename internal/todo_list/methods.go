package todo_list

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-list-api/internal/jwt"
	"todo-list-api/internal/todo_list/task"
	"todo-list-api/pkg/logger"
)

func Get(w http.ResponseWriter, r *http.Request) {
	t, err := jwt.Auth(w, r)
	if err != nil {
		return
	}

	m := jwt.ParseClaims(t)
	id := int(m["sub"].(float64))

	arr, err := GetUserTasks(context.Background(), id)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	resp, err := json.Marshal(arr)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
}

func Create(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.Auth(w, r)
	if err != nil {
		return
	}

	t, err := task.ParseTask(r)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "400 bad request", 400)
		return
	}

	m := jwt.ParseClaims(token)
	userID := int(m["sub"].(float64))
	if userID != t.UserID {
		t.UserID = userID
	}

	err = task.InsertTask(context.Background(), t)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	w.WriteHeader(201)
	return
}

func ParseQuery(r *http.Request) (int, int, error) {
	q := r.URL.Query()

	if len(q["count"]) != 1 || len(q["offset"]) != 1 {
		return 0, 0, fmt.Errorf("bad query")
	}

	q1, q2 := q["count"][0], q["offset"][0]

	count, err := strconv.Atoi(q1)
	if err != nil {
		return 0, 0, fmt.Errorf("bad query")
	}

	offset, err := strconv.Atoi(q2)
	if err != nil {
		return 0, 0, fmt.Errorf("bad query")
	}

	return count, offset, nil
}
