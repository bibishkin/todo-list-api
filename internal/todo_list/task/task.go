package task

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
	"todo-list-api/config"
)

type Task struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Name      string     `json:"name"`
	Body      string     `json:"body"`
	IsDone    bool       `json:"is_done"`
	CreatedAt *time.Time `json:"created_at"`
	Deadline  *time.Time `json:"deadline"`
}

func GetTaskByID(ctx context.Context, pool *pgxpool.Pool, id int) (*Task, error) {
	row := pool.QueryRow(ctx, "SELECT * FROM "+config.DBTask.Table+" WHERE "+config.DBTask.ID+"= $1;", id)
	var (
		userID              int
		name, body          string
		isDone              bool
		createdAt, deadline time.Time
	)
	err := row.Scan(&id, &userID, &name, &body, &isDone, &createdAt, &deadline)
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:        id,
		UserID:    userID,
		Name:      name,
		Body:      body,
		IsDone:    isDone,
		CreatedAt: &createdAt,
		Deadline:  &deadline,
	}, nil
}

func (t *Task) Marshal() ([]byte, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return b, nil
}
