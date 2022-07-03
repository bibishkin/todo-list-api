package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"todo-list-api/pkg/logger"
)

func New(url string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		logger.ErrorLog.Fatalln(err)
	}
	return pool
}
