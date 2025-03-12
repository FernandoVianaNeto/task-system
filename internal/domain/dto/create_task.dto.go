package dto

import "task-system/internal/domain/entities"

type CreateTaskDto struct {
	User    entities.User `json:"user"`
	Title   string        `json:"title"`
	Summary string        `json:"summary"`
}
