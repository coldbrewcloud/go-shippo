package models

// See https://goshippo.com/docs/reference#refunds
type RefundInput struct {
	Transaction string `json:"transaction"`
}

// See https://goshippo.com/docs/reference#refunds
type RefundOutput struct {
	RefundInput
	CommonOutputFields
}
