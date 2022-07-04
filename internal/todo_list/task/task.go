package task

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	. "todo-list-api/config"
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

func GetTaskByID(ctx context.Context, id int) (*Task, error) {
	row := DBPool.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE %s = $1;", DBTask.Table, DBTask.ID), id)
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

func DeleteTaskByID(ctx context.Context, id int) error {
	_, err := DBPool.Exec(ctx, fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", DBTask.Table, DBTask.ID), id)
	return err
}

func InsertTask(ctx context.Context, t *Task) error {
	_, err := DBPool.Exec(
		ctx,
		fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6);", DBTask.Table, DBTask.UserID, DBTask.Name, DBTask.Body, DBTask.IsDone, DBTask.CreatedAt, DBTask.Deadline),
		t.UserID, t.Name, t.Body, t.IsDone, t.CreatedAt, t.Deadline,
	)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) Marshal() ([]byte, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Unmarshal(b []byte) (*Task, error) {
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ParseTask(r *http.Request) (*Task, error) {
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
