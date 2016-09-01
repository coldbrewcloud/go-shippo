package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) GetShippingRates(shipmentObjectID, currencyCode string) ([]*models.RateOutput, error) {
	list := []*models.RateOutput{}
	err := c.doList(http.MethodGet, "/shipments/"+shipmentObjectID+"/rates/"+currencyCode, nil, func(v json.RawMessage) error {
		item := &models.RateOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}

func (c *Client) RetrieveRate(objectID string) (*models.RateOutput, error) {
	output := &models.RateOutput{}
	err := c.do(http.MethodGet, "/rates/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllRates() ([]*models.RateOutput, error) {
	list := []*models.RateOutput{}
	err := c.doList(http.MethodGet, "/rates/", nil, func(v json.RawMessage) error {
		item := &models.RateOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
