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
	ServiceLevelName       string           `json:"servicelevel_name"`
	ServiceLevelToken      string           `json:"servicelevel_token"` // https://goshippo.com/docs/reference#servicelevels
	ServiceLevelTerms      string           `json:"servicelevel_terms"`
	Days                   int              `json:"days"`
	DurationTerms          string           `json:"duration_terms"`
	Trackable              bool             `json:"trackable"`
	Insurance              bool             `json:"insurance"`
	InsuranceAmountLocal   string           `json:"insurance_amount_local"`
	InsuranceCurrencyLocal string           `json:"insurance_currency_local"`
	InsuranceAmount        string           `json:"insurance_amount"`
	InsuranceCurrency      string           `json:"insurance_currency"`
	CarrierAccount         string           `json:"carrier_account"`
	Messages               []*OutputMessage `json:"messages"`
}
