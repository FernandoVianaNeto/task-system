package entities

type User struct {
	Uuid string `json:"uuid"`
	Role string `json:"role"`
	Name string `json:"name"`
}
