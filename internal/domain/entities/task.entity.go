package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Uuid      string    `json:"uuid"`
	Owner     string    `json:"owner"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

func NewTask(
	owner string,
	title string,
	summary string,
) *Task {
	entity := &Task{
		Uuid:    uuid.New().String(),
		Owner:   owner,
		Title:   title,
		Summary: summary,
		Status:  "active",
	}
	return entity
}
