package models

type CarrierAccountInput struct {
	Carrier    string            `json:"carrier"`
	AccountID  string            `json:"account_id"`
	Parameters map[string]string `json:"parameters"`
	Test       bool              `json:"test"`
	Active     bool              `json:"active"`
}

type CarrierAccountOutput struct {
	CarrierAccountInput
	CommonOutputFields
}
