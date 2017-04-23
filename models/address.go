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
	StreetNo      string `json:"street_no,omitempty"`
	Street2       string `json:"street2,omitempty"`
	Street3       string `json:"street3,omitempty"`
	City          string `json:"city,omitempty"`
	Zip           string `json:"zip,omitempty"`
	State         string `json:"state,omtiempty"`
	Country       string `json:"country,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	IsResidential bool   `json:"is_residential"`
	Validate      bool   `json:"validate"`
	Metadata      string `json:"metadata,omitempty"`
}

// See https://goshippo.com/docs/reference#addresses
type Address struct {
	AddressInput
	CommonOutputFields
	IsComplete        bool               `json:"is_complete"`
	Test              bool               `json:"test"`
	ValidationResults *ValidationResults `json:"validation_results"`
}

// See https://goshippo.com/docs/reference#addresses
type ValidationResults struct {
	IsValid  bool             `json:"is_valid"`
	Messages []*OutputMessage `json:"messages"`
}
