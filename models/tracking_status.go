package models

import "time"

const (
	TrackingStatusStatusUnknown   = "UNKNOWN"
	TrackingStatusStatusDelivered = "DELIVERED"
	TrackingStatusStatusTransit   = "TRANSIT"
	TrackingStatusStatusFailure   = "FAILURE"
	TrackingStatusStatusReturned  = "RETURNED"
)

// See https://goshippo.com/docs/reference#tracks
type TrackingStatusInput struct {
	Carrier        string `json:"carrier"`
	TrackingNumber string `json:"tracking_number"`
}

// See https://goshippo.com/docs/reference#tracks
type TrackingStatusOutput struct {
	TrackingStatusInput
	TrackingStatus  *TrackingStatusDict   `json:"tracking_status"`
	TrackingHistory []*TrackingStatusDict `json:"tracking_history"`
}

type TrackingStatusDict struct {
	Status       string                  `json:"status"`
	StatusDetail string                  `json:"status_detail"`
	StatusDate   time.Time               `json:"status_date"`
	Location     *TrackingStatusLocation `json:"location"`
}

type TrackingStatusLocation struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}
