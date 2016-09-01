package models

// See https://goshippo.com/docs/reference#carrier-accounts
type CarrierAccountInput struct {
	Carrier    string            `json:"carrier"`
	AccountID  string            `json:"account_id"`
	Parameters map[string]string `json:"parameters"`
	Test       bool              `json:"test"`
	Active     bool              `json:"active"`
}

// See https://goshippo.com/docs/reference#carrier-accounts
type CarrierAccount struct {
	CarrierAccountInput
	CommonOutputFields
}
