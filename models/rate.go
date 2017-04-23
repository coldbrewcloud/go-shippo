package models

// See https://goshippo.com/docs/reference#rates
type Rate struct {
	CommonOutputFields
	Attributes       []string         `json:"attributes"`
	AmountLocal      string           `json:"amount_local"`
	CurrencyLocal    string           `json:"currency_local"`
	Amount           string           `json:"amount"`
	Currency         string           `json:"currency"`
	Provider         string           `json:"provider"`
	ProviderImage75  string           `json:"provider_image_75"`
	ProviderImage200 string           `json:"provider_image_200"`
	ServiceLevel     *ServiceLevel    `json:"servicelevel"`
	Days             int              `json:"days"`
	DurationTerms    string           `json:"duration_terms"`
	CarrierAccount   string           `json:"carrier_account"`
	Zone             string           `json:"zone"`
	Messages         []*OutputMessage `json:"messages"`
	Test             bool             `json:"test"`
}

type ServiceLevel struct {
	Name  string `json:"name"`
	Token string `json:"token"` // https://goshippo.com/docs/reference#servicelevels
	Terms string `json:"terms"`
}
