package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id        uint           `gorm:"primaryKey"`
	Uuid      string         `json:"uuid" gorm:"type:varchar(255);not null"`
	UserUuid  string         `json:"user_uuid" gorm:"type:varchar(255);not null"`
	Title     string         `json:"title" gorm:"type:varchar(255);not null"`
	Summary   string         `json:"summary" gorm:"type:varchar(500);not null"`
	Status    string         `json:"status" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
