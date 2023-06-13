package request

type UpdateStatusRequest struct {
	LastId string `json:"lastId"`
	Status string `json:"status"`
}
