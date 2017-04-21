package models

import "time"

// See https://goshippo.com/docs/reference#manifests
type ManifestInput struct {
	CarrierAccount string    `json:"carrier_account,omitempty"`
	ShipmentDate   time.Time `json:"shipment_date,omitempty"`
	AddressFrom    string    `json:"address_from,omitempty"`
	Transactions   []string  `json:"transactions,omitempty"`
	Documents      []string  `json:"documents,omitempty"`
	Async          bool      `json:"async"`
}

// See https://goshippo.com/docs/reference#manifests
type Manifest struct {
	ManifestInput
	CommonOutputFields
	State string `json:"state,omitempty"`
}
