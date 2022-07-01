package main

const (
	certPath = "cert.cer"
	keyPath  = "pkey.pkey"
	addr     = "192.168.100.12"

	databaseConnString = "postgres://postgres:1111@127.0.0.1:5432/todolist"
)

var dbConn = NewConnect(databaseConnString)
