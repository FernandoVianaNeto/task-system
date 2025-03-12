package entities

import "time"

type Task struct {
	Id        int    `json:"id"`
	Uuid      string `json:"uuid"`
	User      User   `json:"user"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

func NewTask(
	user User,
	title string,
	summary string,
) *Task {
	entity := &Task{
		Id:        1,
		Uuid:      "",
		User:      user,
		Title:     title,
		Summary:   summary,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.000Z"),
		Status:    "active",
	}
	return entity
}
