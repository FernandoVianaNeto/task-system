package dto

type CreateUserDto struct {
	Role string `json:"role"`
	Name string `json:"name"`
}
