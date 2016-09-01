package client

import (
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) GetTrackingUpdate(carrier, trackingNumber string) (*models.TrackingStatusOutput, error) {
	output := &models.TrackingStatusOutput{}
	err := c.do(http.MethodGet, "/trakcs/"+carrier+"/"+trackingNumber, nil, output)
	return output, err
}

func (c *Client) RegisterTrackingWebhook(carrier, trackingNumber string) (*models.TrackingStatusOutput, error) {
	output := &models.TrackingStatusOutput{}
	err := c.do(http.MethodPost, "/trakcs/"+carrier+"/"+trackingNumber, nil, output)
	return output, err
}
