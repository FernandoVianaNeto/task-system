package entities

import "github.com/google/uuid"

type User struct {
	Uuid     string `json:"uuid"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUser(
	role string,
	name string,
	email string,
	password string,
) *User {
	entity := &User{
		Uuid:     uuid.New().String(),
		Role:     role,
		Name:     name,
		Email:    email,
		Password: password,
	}
	return entity
}
