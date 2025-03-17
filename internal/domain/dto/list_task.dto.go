package dto

type ListTaskDto struct {
	Uuid  *string `json:"uuid"`
	Owner *string `json:"owner"`
}
