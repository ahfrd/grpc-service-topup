package entity

type FeeTopUpEntity struct {
	Id     string `json:"id"`
	Method string `json:"method"`
	Fee    int    `json:"fee"`
}