package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/d5/go-shippo/models"
)

// CreateAddress creates a new address object.
func (c *Client) CreateAddress(input *models.AddressInput) (*models.AddressOutput, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.AddressOutput{}
	err := c.do(http.MethodPost, "/addresses/", input, output)
	return output, err
}

// RetrieveAddress retrieves an existing address by object id.
func (c *Client) RetrieveAddress(objectID string) (*models.AddressOutput, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.AddressOutput{}
	err := c.do(http.MethodGet, "/addresses/"+objectID, nil, output)
	return output, err
}

// ListAllAddresses lists all addresses.
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
