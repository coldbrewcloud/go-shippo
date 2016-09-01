package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/d5/go-shippo/models"
)

// GetShippingRates gets rates for a shipping object.
func (c *Client) GetShippingRates(shipmentObjectID, currencyCode string) ([]*models.RateOutput, error) {
	if shipmentObjectID == "" {
		return nil, errors.New("Empty shipment object ID")
	}
	if currencyCode == "" {
		return nil, errors.New("Empty currency code")
	}

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

// RetrieveRate retrieves an existing rate by object id.
func (c *Client) RetrieveRate(objectID string) (*models.RateOutput, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.RateOutput{}
	err := c.do(http.MethodGet, "/rates/"+objectID, nil, output)
	return output, err
}

// ListAllRates lists all the rate objects that have been previously obtained through GetShippingRates function.
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
