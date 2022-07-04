package config

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"os"
	"todo-list-api/internal/db"
	"todo-list-api/pkg/logger"
)

var Addr string

var JWT *jwt

var DBPool *pgxpool.Pool

var DBUser *dbUser

var DBTask *dbTask

type dbUser struct {
	Table    string
	ID       string
	Email    string
	Password string
}

type dbTask struct {
	Table     string
	ID        string
	UserID    string
	Name      string
	Body      string
	IsDone    string
	CreatedAt string
	Deadline  string
}

type jwt struct {
	Cookie string
	Key    []byte
}

func Init() {

	if err := godotenv.Load(); err != nil {
		logger.InfoLog.Println("No .env file found")
	}

	Addr = os.Getenv("ADDR")
	JWT = &jwt{
		Cookie: "jwt",
		Key:    []byte(os.Getenv("JWT_KEY")),
	}

	DBPool = db.New(os.Getenv("DB_URL"))

	DBUser = &dbUser{
		Table:    "users",
		ID:       "id",
		Email:    "email",
		Password: "password",
	}

	DBTask = &dbTask{
		Table:     "todo_list",
		ID:        "id",
		UserID:    "user_id",
		Name:      "task_name",
		Body:      "task_body",
		IsDone:    "is_done",
		CreatedAt: "created_at",
		Deadline:  "deadline",
	}

}
