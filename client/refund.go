package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

// CreateRefund creates a new refund object.
func (c *Client) CreateRefund(input *models.RefundInput) (*models.RefundOutput, error) {
	output := &models.RefundOutput{}
	err := c.do(http.MethodPost, "/refunds/", input, output)
	return output, err
}

// RetrieveRefund retrieves an existing refund by object id.
func (c *Client) RetrieveRefund(objectID string) (*models.RefundOutput, error) {
	output := &models.RefundOutput{}
	err := c.do(http.MethodGet, "/refunds/"+objectID, nil, output)
	return output, err
}

// ListAllRefunds list all refund objects.
func (c *Client) ListAllRefunds() ([]*models.RefundOutput, error) {
	list := []*models.RefundOutput{}
	err := c.doList(http.MethodGet, "/refunds/", nil, func(v json.RawMessage) error {
		item := &models.RefundOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
