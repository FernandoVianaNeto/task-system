package dto

type GetUserByUuidDto struct {
	Uuid string `json:"uuid"`
}

type GetUserByEmailDto struct {
	Email string `json:"email"`
}
