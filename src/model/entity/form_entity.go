package entity

type FormStructEntity struct {
	Id      string           `json:"id"`
	Method  string           `json:"method"`
	Details []BankNameEntity `json:"details"`
}

type BankNameEntity struct {
	Code     string `json:"code"`
	BankName string `json:"bankName"`
}
