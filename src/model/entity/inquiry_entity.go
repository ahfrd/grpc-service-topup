package entity

type InquiryEntity struct {
	Method          string           `json:"method"`
	Nominal         string           `json:"nominalTopUp"`
	Fee             string           `json:"feeTopUp"`
	TransactionName string           `json:"transactionName"`
	PhoneNumb       string           `json:"phoneNumber"`
	Card            DebitCardRequest `json:"cardInfo"`
	CodeNumb        string           `json:"codeNumber"`
}

type DebitCardRequest struct {
	CardNumb    string `json:"nomorKartu"`
	Valid       string `json:"validUntil"`
	SecurityNum string `json:"securityNum"`
}
