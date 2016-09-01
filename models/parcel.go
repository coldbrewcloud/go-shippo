package models

const (
	ParcelDistanceUnitCentiMeter = "cm"
	ParcelDistanceUnitInch       = "in"
	ParcelDistanceUnitFoot       = "ft"
	ParcelDistanceUnitMilliMeter = "mm"
	ParcelDistanceUnitMeter      = "m"
	ParcelDistanceUnitYard       = "yd"

	ParcelMassUnitGram     = "g"
	ParcelMassUnitOunce    = "oz"
	ParcelMassUnitPound    = "lb"
	ParcelMassUnitKiloGram = "kg"
)

type ParcelInput struct {
	Length       string       `json:"length"`
	Width        string       `json:"width"`
	Height       string       `json:"height"`
	DistanceUnit string       `json:"distance_unit"`
	Weight       string       `json:"weight"`
	MassUnit     string       `json:"mass_unit"`
	Template     string       `json:"template,omitempty"` // https://goshippo.com/docs/reference#parceltemplates
	Metadata     string       `json:"metadata"`
	Extra        *ParcelExtra `json:"extra"`
}

type ParcelExtra struct {
	COD       *COD             `json:"COD"`
	Insurance *ParcelInsurance `json:"insurance"`
}

type ParcelInsurance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Provider string `json:"provider"`
	Content  string `json:"content"`
}

type ParcelOutput struct {
	ParcelInput
	CommonOutputFields
}
