package user

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"todo-list-api/config"
)

type User struct {
	ID       int    `json:"ID"`
	Email    string `json:"Email"`
	Password string `json:"password"`
}

func New(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func InsertUser(ctx context.Context, pool *pgxpool.Pool, u *User) error {
	_, err := pool.Exec(
		ctx,
		"INSERT INTO "+config.DBUser.Table+"("+config.DBUser.Email+","+config.DBUser.Password+")"+" VALUES ($1, $2);",
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(ctx context.Context, pool *pgxpool.Pool, id int) (*User, error) {

	row := pool.QueryRow(ctx, "SELECT * FROM "+config.DBUser.Table+" WHERE "+config.DBUser.ID+"= $1;", id)
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

func GetUserByEmail(ctx context.Context, pool *pgxpool.Pool, email string) (*User, error) {
	row := pool.QueryRow(ctx, "SELECT * FROM "+config.DBUser.Table+" WHERE "+config.DBUser.Email+"= $1;", email)
	var id int
	var password string
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
