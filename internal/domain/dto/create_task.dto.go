package dto

type CreateTaskDto struct {
	UserUuid string `json:"user_uuid"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
}
