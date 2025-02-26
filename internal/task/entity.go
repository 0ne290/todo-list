package task

import (
	"errors"
	"time"
)

type Status string

const (
	New        Status = "new"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Task struct {
	Id          int `json:"id"`
	Title       string `json:"title"`
	Description *string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(title string, description *string, createdAt time.Time) *Task {
	return &Task{0, title, description, New, createdAt, createdAt}
}

func (task *Task) Update() error {
	switch task.Status {

	case New:
		task.Status = InProgress

		return nil

	case InProgress:
		task.Status = Done

		return nil

	default:
		return errors.New("status is invalid")
	}
}
