package models

import "time"

// See https://goshippo.com/docs/reference#transactions
type TransactionInput struct {
	Rate          string `json:"rate,omitempty"`
	Metadata      string `json:"metadata,omitempty"`
	LabelFileType string `json:"label_file_type,omitempty"`
	Async         bool   `json:"async"`

	Shipment         *ShipmentInput `json:"shipment,omitempty"`              // instant call only: https://goshippo.com/docs/reference#transactions-create-instant
	CarrierAccount   string         `json:"carrier_account,omitempty"`       // instant call only: https://goshippo.com/docs/reference#transactions-create-instant
	ServerLevelToken string         `json:"servericelevel_token, omitempty"` // instant call only: https://goshippo.com/docs/reference#transactions-create-instant
}

// See https://goshippo.com/docs/reference#transactions
type Transaction struct {
	TransactionInput
	CommonOutputFields
	ObjectState          string           `json:"object_state,omitempty"`
	Status               string           `json:"status,omitempty"`
	Test                 bool             `json:"test"`
	TrackingNumber       string           `json:"tracking_number,omitempty"`
	TrackingStatus       string           `json:"tracking_status,omitempty"`
	TrackingURLProvider  string           `json:"tracking_url_provider,omitempty"`
	Eta                  time.Time        `json:"eta,omitempty"`
	LabelURL             string           `json:"label_url,omitempty"`
	CommercialInvoiceURL string           `json:"commercial_invoice_url,omitempty"`
	Messages             []*OutputMessage `json:"messages,omitempty"`
	QRCodeURL            string           `json:"qr_code_url,omitempty"`
	Async                bool             `json:"async"`
}
