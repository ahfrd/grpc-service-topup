package entity

type UpdateBallance struct {
	PhoneNumber         string `json:"phoneNumber"`
	CategoryTransaction string `json:"categoryTransaction"`
	NominalTransaction  string `json:"nominalTransaction"`
	SecurityCode        string `json:"securityCode"`
	FeeTransaction      string `json:"feeTransaction"`
}
