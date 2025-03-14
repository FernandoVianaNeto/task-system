package entities

import "time"

type Task struct {
	Id        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	User      User      `json:"user"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

func NewTask(
	user User,
	title string,
	summary string,
) *Task {
	entity := &Task{
		Id:      1,
		Uuid:    "",
		User:    user,
		Title:   title,
		Summary: summary,
		Status:  "active",
	}
	return entity
}
