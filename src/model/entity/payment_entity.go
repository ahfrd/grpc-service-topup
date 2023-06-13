package entity

type InsertHistoryTopUpEntity struct {
	ReffNum          string `json:"refferenceNumber"`
	Nominal          string `json:"nominalTopUp"`
	Fee              string `json:"feeTopUp"`
	TopUpDestination string `json:"topUpDestination"`
	TopUpSource      string `json:"topUpSource"`
	HistoryId        string `json:"historyId"`
}
