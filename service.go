package main

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"io"
)

func GetUser(body io.Reader) (*User, error) {

	readBody, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	user := User{}

	err = json.Unmarshal(readBody, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func IsUserExist(u *User) (bool, error) {
	dbConn.Mutex.Lock()
	row := dbConn.Conn.QueryRow(context.Background(), "SELECT FROM users WHERE name = $1;", u.Name)
	dbConn.Mutex.Unlock()

	err := row.Scan()

	switch err {
	case pgx.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func AddUser(u *User) error {
	dbConn.Mutex.Lock()
	_, err := dbConn.Conn.Exec(context.Background(), "INSERT INTO users (name, pass) VALUES ($1, $2);", u.Name, u.Password)
	dbConn.Mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}
