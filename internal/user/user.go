package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	. "todo-list-api/config"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func InsertUser(ctx context.Context, u *User) error {
	_, err := DBPool.Exec(
		ctx,
		fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2);", DBUser.Table, DBUser.Email, DBUser.Password),
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(ctx context.Context, id int) (*User, error) {

	row := DBPool.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = $1;", DBUser.Table, DBUser.ID))
	var email, password string
	err := row.Scan(&id, &email, &password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}, nil
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := DBPool.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = $1;", DBUser.Table, DBUser.Email), email)
	var (
		id       int
		password string
	)
	err := row.Scan(&id, &email, &password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) Marshal() ([]byte, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Unmarshal(b []byte) (*User, error) {
	u := User{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func ParseUser(r *http.Request) (*User, error) {
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
