package models

import "encoding/json"

// See https://goshippo.com/docs/reference#batches
type BatchInput struct {
	DefaultCarrierAccount    string                `json:"default_carrier_account,omitempty"`
	DefaultServiceLevelToken string                `json:"default_servicelevel_token,omitempty"`
	LabelFileType            string                `json:"label_filetype,omitempty"`
	Metadata                 string                `json:"metadata,omitempty"`
	BatchShipments           []*BatchShipmentInput `json:"batch_shipments"`
}

// See https://goshippo.com/docs/reference#batches
type Batch struct {
	//BatchInput // seems that output format is slightly different
	CommonOutputFields
	Status                   string              `json:"status,omitempty"`
	DefaultCarrierAccount    string              `json:"default_carrier_account,omitempty"`
	DefaultServiceLevelToken string              `json:"default_servicelevel_token,omitempty"`
	LabelFileType            string              `json:"label_filetype,omitempty"`
	Metadata                 string              `json:"metadata,omitempty"`
	BatchShipments           *BatchShipmentList  `json:"batch_shipments"`
	LabelURL                 []string            `json:"label_url"`
	ObjectResults            *BatchObjectResults `json:"object_results"`
}

type BatchShipmentList struct {
	Count    int              `json:"count"`
	Next     string           `json:"next,omitempty"`
	Previous string           `json:"previous,omitempty"`
	Results  []*BatchShipment `json:"results"`
}

type BatchObjectResults struct {
	PurchaseSucceeded int `json:"purchase_succeeded"`
	PurchaseFailed    int `json:"purchase_failed"`
	CreationFailed    int `json:"creation_failed"`
	CreationSucceeded int `json:"creation_succeeded"`
}

// See https://goshippo.com/docs/reference#batches-batchshipments
type BatchShipmentInput struct {
	CarrierAccount    string `json:"carrier_account,omitempty"`
	ServiceLevelToken string `json:"servicelevel_token,omitempty"`
	Shipment          string `json:"shipment"`
	Metadata          string `json:"metadata,omitempty"`
}

// See https://goshippo.com/docs/reference#batches-batchshipments
type BatchShipment struct {
	BatchShipmentInput
	CommonOutputFields
	Transaction string            `json:"transaction"`
	Messages    []json.RawMessage `json:"messages"` // cannot find format in documentation
}
