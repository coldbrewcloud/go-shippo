package models

type RefundInput struct {
	Transaction string `json:"transaction"`
}

type RefundOutput struct {
	RefundInput
	CommonOutputFields
}
