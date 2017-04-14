package models

import "time"

// See https://goshippo.com/docs/reference#manifests
type ManifestInput struct {
	CarrierAccount string    `json:"carrier_account"`
	ShipmentDate   time.Time `json:"shipment_date"`
	AddressFrom    string    `json:"address_from"`
	Transactions   []string  `json:"transactions"`
	Async          bool      `json:"async"`
}

// See https://goshippo.com/docs/reference#manifests
type Manifest struct {
	ManifestInput
	CommonOutputFields
}
