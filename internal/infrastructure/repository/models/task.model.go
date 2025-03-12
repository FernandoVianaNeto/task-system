package models

import "task-system/internal/domain/entities"

type TaskModel struct {
	Id        int           `json:"id" gorm:"column:primaryKey"`
	Uuid      string        `json:"uuid" gorm:"column:uuid"`
	User      entities.User `json:"user" gorm:"column:user"`
	Title     string        `json:"title" gorm:"column:title"`
	Summary   string        `json:"summary" gorm:"column:summary"`
	CreatedAt string        `json:"created_at" gorm:"column:created_at"`
	Status    string        `json:"status" gorm:"column:status"`
}
