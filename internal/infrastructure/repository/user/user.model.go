package repository_user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey"`
	Uuid      string         `json:"uuid" gorm:"type:varchar(255);not null"`
	Role      string         `json:"role" gorm:"type:varchar(255);not null"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Password  string         `json:"password" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
