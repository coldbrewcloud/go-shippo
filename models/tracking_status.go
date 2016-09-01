package models

type TrackingStatusInput struct{}

type TrackingStatusOutput struct {
	TrackingStatusInput
}

type TrackingStatusDict struct {
	Status string `json:"status"`
}
