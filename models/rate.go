package models

// See https://goshippo.com/docs/reference#rates
type Rate struct {
	CommonOutputFields
	Attributes             []string         `json:"attributes"`
	AmountLocal            string           `json:"amount_local"`
	CurrencyLocal          string           `json:"currency_local"`
	Amount                 string           `json:"amount"`
	Currency               string           `json:"currency"`
	Provider               string           `json:"provider"`
	ProviderImage75        string           `json:"providerImage75"`
	ProviderImage200       string           `json:"providerImage200"`
	ServiceLevel           *ServiceLevel    `json:"serviceLevel"`
	Days                   int              `json:"days"`
	DurationTerms          string           `json:"duration_terms"`
	Insurance              bool             `json:"insurance"`
	InsuranceAmountLocal   string           `json:"insurance_amount_local"`
	InsuranceCurrencyLocal string           `json:"insurance_currency_local"`
	InsuranceAmount        string           `json:"insurance_amount"`
	InsuranceCurrency      string           `json:"insurance_currency"`
	CarrierAccount         string           `json:"carrier_account"`
	Messages               []*OutputMessage `json:"messages"`
}

type ServiceLevel struct {
	Name  string `json:"name"`
	Token string `json:"token"` // https://goshippo.com/docs/reference#servicelevels
	Terms string `json:"terms"`
}
