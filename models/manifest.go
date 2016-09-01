package models

import "time"

type ManifestInput struct {
	CarrierAccount string    `json:"carrier_account"`
	SubmissionDate time.Time `json:"submission_date"`
	AddressFrom    string    `json:"address_from"`
	Transactions   []string  `json:"transactions"`
	Async          bool      `json:"async"`
}

type ManifestOutput struct {
	ManifestInput
	CommonOutputFields
}
