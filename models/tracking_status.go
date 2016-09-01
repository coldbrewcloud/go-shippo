package models

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
	TrackingStatus  *TrackingStatusDict   `json:"tracking_status,omitempty"`
	TrackingHistory []*TrackingStatusDict `json:"tracking_history,omitempty"`
}

type TrackingStatusDict struct {
	Status       string                  `json:"status,omitempty"`
	StatusDetail string                  `json:"status_detail,omitempty"`
	StatusDate   string                  `json:"status_date,omitempty"`
	Location     *TrackingStatusLocation `json:"location,omitempty"`
}

type TrackingStatusLocation struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
}
