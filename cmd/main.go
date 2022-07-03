package main

import (
	"todo-list-api/config"
	"todo-list-api/internal/router"
	"todo-list-api/pkg/server"
)

func main() {
	config.Init()
	srv := server.NewServer(config.Addr, router.NewRouter())
	srv.Run()
}
