package dto

type UpdateTaskStatusDto struct {
	TaskUuid   string `json:"uuid"`
	TaskStatus string `json:"status"`
	UserUuid   string `json:"owner"`
}
