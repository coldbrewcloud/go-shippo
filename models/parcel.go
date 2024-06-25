package models

// See https://goshippo.com/docs/reference#parcels
type ParcelInput struct {
	Length       string       `json:"length"`
	Width        string       `json:"width"`
	Height       string       `json:"height"`
	DistanceUnit string       `json:"distance_unit"`
	Weight       string       `json:"weight"`
	MassUnit     string       `json:"mass_unit"`
	Template     string       `json:"template,omitempty"` // https://goshippo.com/docs/reference#parceltemplates
	Extra        *ParcelExtra `json:"extra,omitempty"`
	Metadata     string       `json:"metadata,omitempty"`
}

type ParcelExtra struct {
	COD        *ParcelCOD       `json:"COD,omitempty"`
	Insurance  *ParcelInsurance `json:"insurance,omitempty"`
	Reference1 string           `json:"reference_1,omitempty"`
	Reference2 string           `json:"reference_2,omitempty"`
}

type ParcelCOD struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
}

type ParcelInsurance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Provider string `json:"provider,omitempty"`
	Content  string `json:"content"`
}

// See https://goshippo.com/docs/reference#parcels
type Parcel struct {
	ParcelInput
	CommonOutputFields
	ObjectState string `json:"object_state,omitempty"`
	Test        bool   `json:"test"`
}
