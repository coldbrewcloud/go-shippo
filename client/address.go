package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) CreateAddress(input *models.AddressInput) (*models.AddressOutput, error) {
	output := &models.AddressOutput{}
	err := c.do(http.MethodPost, "/addresses/", input, output)
	return output, err
}

func (c *Client) RetrieveAddress(objectID string) (*models.AddressOutput, error) {
	output := &models.AddressOutput{}
	err := c.do(http.MethodGet, "/addresses/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllAddresses() ([]*models.AddressOutput, error) {
	list := []*models.AddressOutput{}
	err := c.doList(http.MethodGet, "/addresses/", nil, func(v json.RawMessage) error {
		item := &models.AddressOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
