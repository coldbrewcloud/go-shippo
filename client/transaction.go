package client

import (
	"net/http"

	"encoding/json"

	"github.com/d5/go-shippo/models"
)

func (c *Client) PurchaseShippingLabel(input *models.TransactionInput) (*models.TransactionOutput, error) {
	output := &models.TransactionOutput{}
	err := c.do(http.MethodPost, "/transactions/", input, output)
	return output, err
}

func (c *Client) RetrieveTransaction(objectID string) (*models.TransactionOutput, error) {
	output := &models.TransactionOutput{}
	err := c.do(http.MethodGet, "/transactions/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllTransactions() ([]*models.TransactionOutput, error) {
	list := []*models.TransactionOutput{}
	err := c.doList(http.MethodGet, "/transactions/", nil, func(v json.RawMessage) error {
		item := &models.TransactionOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
