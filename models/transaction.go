package models

// See https://goshippo.com/docs/reference#transactions
type TransactionInput struct {
	Rate          *Rate  `json:"rate"`
	Metadata      string `json:"metadata,omitempty"`
	LabelFileType string `json:"label_file_type"`
	Async         bool   `json:"async"`
}

// See https://goshippo.com/docs/reference#transactions
type Transaction struct {
	TransactionInput
	CommonOutputFields
	ObjectStatus         string                `json:"object_status,omitempty"`
	WasTest              bool                  `json:"was_test"`
	TrackingNumber       string                `json:"tracking_number,omitempty"`
	TrackingStatus       *TrackingStatusDict   `json:"tracking_status,omitempty"`
	TrackingHistory      []*TrackingStatusDict `json:"tracking_history,omitempty"`
	TrackingURLProvider  string                `json:"tracking_url_provider,omitempty"`
	LabelURL             string                `json:"label_url,omitempty"`
	CommercialInvoiceURL string                `json:"commercial_invoice_url,omitempty"`
	Messages             []*OutputMessage      `json:"messages,omitempty"`
	Async                bool                  `json:"async"`
}
