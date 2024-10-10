package model

import (
	"fmt"
	"time"
)

type Status string

const (
	TODO       Status = "todo"
	InProgress Status = "in_progress"
	DONE       Status = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func TaskFormatter(task Task) string {
	return fmt.Sprintf(
		"task id: %d, task description: %s, status: %s\nupdated_at: %s, created_at: %s\n%s\n",
		task.Id,
		task.Description,
		task.Status,
		task.UpdatedAt,
		task.CreatedAt,
		"--------",
	)
}
