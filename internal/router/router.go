package router

import (
	"github.com/gorilla/mux"
	"todo-list-api/internal/todo_list"
	"todo-list-api/internal/todo_list/task"
	"todo-list-api/internal/user"
)

func NewRouter() *mux.Router {
	m := mux.NewRouter()

	m.HandleFunc("/api/user/signup", user.SignUp).Methods("POST")
	m.HandleFunc("/api/user/signin", user.SignIn).Methods("POST")
	m.HandleFunc("/api/todo", todo_list.Get).Methods("GET")
	m.HandleFunc("/api/todo", todo_list.Create).Methods("POST")
	m.HandleFunc("/api/todo/{id:[0-9]+}", task.Get).Methods("GET")
	m.HandleFunc("/api/todo/{id:[0-9]+}", task.Update).Methods("PUT", "PATCH")
	m.HandleFunc("/api/todo/{id:[0-9]+}", task.Delete).Methods("DELETE")

	return m
}
