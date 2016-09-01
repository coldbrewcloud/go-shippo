package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/d5/go-shippo/models"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) PurchaseShippingLabel(input *models.TransactionInput) (*models.TransactionOutput, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.TransactionOutput{}
	err := c.do(http.MethodPost, "/transactions/", input, output)
	return output, err
}

// RetrieveTransaction retrieves an existing transaction by object id.
func (c *Client) RetrieveTransaction(objectID string) (*models.TransactionOutput, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.TransactionOutput{}
	err := c.do(http.MethodGet, "/transactions/"+objectID, nil, output)
	return output, err
}

// ListAllTransactions lists all transaction objects.
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
