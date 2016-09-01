package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) CreateShipment(input *models.ShipmentInput) (*models.ShipmentOutput, error) {
	output := &models.ShipmentOutput{}
	err := c.do(http.MethodPost, "/shipments/", input, output)
	return output, err
}

func (c *Client) RetrieveShipment(objectID string) (*models.ShipmentOutput, error) {
	output := &models.ShipmentOutput{}
	err := c.do(http.MethodGet, "/shipments/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllShipments() ([]*models.ShipmentOutput, error) {
	list := []*models.ShipmentOutput{}
	err := c.doList(http.MethodGet, "/shipments/", nil, func(v json.RawMessage) error {
		item := &models.ShipmentOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
