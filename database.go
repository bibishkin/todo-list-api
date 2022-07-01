package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"sync"
)

type Conn struct {
	Conn  *pgx.Conn
	Mutex sync.Mutex
}

func NewConnect(connStr string) *Conn {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	return &Conn{
		Conn:  conn,
		Mutex: sync.Mutex{},
	}
}
