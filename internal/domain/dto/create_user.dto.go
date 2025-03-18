package dto

type CreateUserDto struct {
	Role     string `json:"role"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
