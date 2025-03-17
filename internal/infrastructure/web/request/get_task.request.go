package infra_request

type ListTaskRequest struct {
	Uuid  string `uri:"uuid"`
	Owner string `uri:"owner"`
}
