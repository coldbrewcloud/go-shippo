package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) CreateCarrierAccount(input *models.CarrierAccountInput) (*models.CarrierAccountOutput, error) {
	output := &models.CarrierAccountOutput{}
	err := c.do(http.MethodPost, "/carrier_accounts/", input, output)
	return output, err
}

func (c *Client) RetrieveCarrierAccount(objectID string) (*models.CarrierAccountOutput, error) {
	output := &models.CarrierAccountOutput{}
	err := c.do(http.MethodGet, "/carrier_accounts/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllCarrierAccounts() ([]*models.CarrierAccountOutput, error) {
	list := []*models.CarrierAccountOutput{}
	err := c.doList(http.MethodGet, "/carrier_accounts/", nil, func(v json.RawMessage) error {
		item := &models.CarrierAccountOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}

func (c *Client) UpdateCarrierAccount(objectID string, input *models.CarrierAccountInput) (*models.CarrierAccountOutput, error) {
	output := &models.CarrierAccountOutput{}
	err := c.do(http.MethodPut, "/carrier_accounts/"+objectID, input, output)
	return output, err
}
