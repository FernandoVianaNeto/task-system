package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	UserUuid  string    `json:"user_uuid"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

func NewTask(
	userUuid string,
	title string,
	summary string,
) *Task {
	entity := &Task{
		Id:       1,
		Uuid:     uuid.New().String(),
		UserUuid: userUuid,
		Title:    title,
		Summary:  summary,
		Status:   "active",
	}
	return entity
}
