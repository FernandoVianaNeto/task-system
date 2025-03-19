package infra_request

type CreateTaskRequest struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}
