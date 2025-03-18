package dto

type ListTaskDto struct {
	Uuid     *string `json:"uuid"`
	Owner    *string `json:"owner"`
	UserUuid string  `json:"userUuid"`
	Role     string  `json:"role"`
}
