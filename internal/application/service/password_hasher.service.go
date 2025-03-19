package service

import (
	domain_service "task-system/internal/domain/service"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasherService struct{}

func NewPasswordHasherService() domain_service.PasswordHasherServiceInterface {
	return &PasswordHasherService{}
}

func (p *PasswordHasherService) HashPassword(password string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}
