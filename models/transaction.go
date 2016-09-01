package models

const (
	TransactionLabelFileTypePNG    = "PNG"
	TransactionLabelFileTypePDF    = "PDF"
	TransactionLabelFileTypePDF4X6 = "PDF_4X6"
	TransactionLabelFileTypeZPLII  = "ZPLII"
)

type TransactionInput struct {
	Rate          string `json:"rate"`
	Metadata      string `json:"metadata,omitempty"`
	LabelFileType string `json:"label_file_type"`
	Async         bool   `json:"async"`
}

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
