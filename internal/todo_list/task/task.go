package task

import "time"

type Task struct {
	ID       int
	UserID   int
	Name     string
	Body     string
	Done     bool
	Created  *time.Time
	Deadline *time.Time
}
