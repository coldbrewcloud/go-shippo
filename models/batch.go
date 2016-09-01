package models

import "encoding/json"

type BatchInput struct {
	DefaultCarrierAccount    string                `json:"default_carrier_account,omitempty"`
	DefaultServiceLevelToken string                `json:"default_servicelevel_token,omitempty"`
	LabelFileType            string                `json:"label_filetype,omitempty"`
	Metadata                 string                `json:"metadata,omitempty"`
	BatchShipments           []*BatchShipmentInput `json:"batch_shipments"`
}

type BatchOutput struct {
	//BatchInput // seems that output format is slightly different
	CommonOutputFields
	DefaultCarrierAccount    string                   `json:"default_carrier_account,omitempty"`
	DefaultServiceLevelToken string                   `json:"default_servicelevel_token,omitempty"`
	LabelFileType            string                   `json:"label_filetype,omitempty"`
	Metadata                 string                   `json:"metadata,omitempty"`
	BatchShipments           *BatchOutputShipmentList `json:"batch_shipments"`
	LabelURL                 []string                 `json:"label_url"`
	ObjectResults            *BatchObjectResults      `json:"object_results"`
}

type BatchOutputShipmentList struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next,omitempty"`
	Previous string                 `json:"previous,omitempty"`
	Results  []*BatchShipmentOutput `json:"results"`
}

type BatchObjectResults struct {
	PurchaseSucceeded int `json:"purchase_succeeded"`
	PurchaseFailed    int `json:"purchase_failed"`
	CreationFailed    int `json:"creation_failed"`
	CreationSucceeded int `json:"creation_succeeded"`
}

type BatchShipmentInput struct {
	CarrierAccount    string `json:"carrier_account,omitempty"`
	ServiceLevelToken string `json:"servicelevel_token,omitempty"`
	Shipment          string `json:"shipment"`
	Metadata          string `json:"metadata,omitempty"`
}

type BatchShipmentOutput struct {
	BatchShipmentInput
	CommonOutputFields
	Transaction string            `json:"transaction"`
	Messages    []json.RawMessage `json:"messages"` // cannot find format in documentation
}
