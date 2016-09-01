package models

// See https://goshippo.com/docs/reference#transactions
type TransactionInput struct {
	Rate          string `json:"rate"`
	Metadata      string `json:"metadata,omitempty"`
	LabelFileType string `json:"label_file_type"`
	Async         bool   `json:"async"`
}

// See https://goshippo.com/docs/reference#transactions
type TransactionOutput struct {
	TransactionInput
	CommonOutputFields
	ObjectStatus         string                `json:"object_status"`
	WasTest              bool                  `json:"was_test"`
	TrackingNumber       string                `json:"tracking_number"`
	TrackingStatus       *TrackingStatusDict   `json:"tracking_status"`
	TrackingHistory      []*TrackingStatusDict `json:"tracking_history"`
	TrackingURLProvider  string                `json:"tracking_url_provider"`
	LabelURL             string                `json:"label_url"`
	CommercialInvoiceURL string                `json:"commercial_invoice_url"`
	Messages             []OutputMessage       `json:"messages"`
	Async                bool                  `json:"async"`
}
