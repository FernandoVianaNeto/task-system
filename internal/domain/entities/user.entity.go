package entities

import "github.com/google/uuid"

type User struct {
	Uuid string `json:"uuid"`
	Role string `json:"role"`
	Name string `json:"name"`
}

func NewUser(
	role string,
	name string,
) *User {
	entity := &User{
		Uuid: uuid.New().String(),
		Role: role,
		Name: name,
	}
	return entity
}
