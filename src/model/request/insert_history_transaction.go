package request

type InsertHistoryTransaction struct {
	WalletId            string `json:"walletId"`
	TransactionName     string `json:"transactionName"`
	NominalTransaction  string `json:"nominalTransaction"`
	FeeTransaction      string `json:"feeTransaction"`
	CategoryTransaction string `json:"categoryTransaction"`
	Status              string `json:"status"`
	TransactionService  string `json:"transactionService"`
	RefferenceNumber    string `json:"refferenceNumber"`
}
