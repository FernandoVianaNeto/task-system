package infra_request

type UpdateTaskStatusRequest struct {
	TaskUuid  string `json:"task_uuid"`
	NewStatus string `json:"status"`
}
