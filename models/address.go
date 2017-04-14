package models

const (
	AddressObjectSourceFullyEntered     = "FULLY_ENTERED"
	AddressObjectSourcePartiallyEntered = "PARTIALLY_ENTERED"
	AddressObjectSourceValidator        = "VALIDATOR"
)

// See https://goshippo.com/docs/reference#addresses
type AddressInput struct {
	Name          string `json:"name,omitempty"`
	Company       string `json:"company,omitempty"`
	Street1       string `json:"street1,omitempty"`
	Street2       string `json:"street2,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omtiempty"`
	Zip           string `json:"zip,omitempty"`
	Country       string `json:"country"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	IsResidential bool   `json:"is_residential"`
	Validate      bool   `json:"validate"`
	Metadata      string `json:"metadata,omitempty"`
}

// See https://goshippo.com/docs/reference#addresses
type Address struct {
	IsComplete bool `json:"is_complete"`
	AddressInput
	CommonOutputFields
	ValidateResults *ValidateResults `json:"messages"`
}

// See https://goshippo.com/docs/reference#addresses
type ValidateResults struct {
	IsValid  bool             `json:"is_valid"`
	Messages []*OutputMessage `json:"messages`
}
