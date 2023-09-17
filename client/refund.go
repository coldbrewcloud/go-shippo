package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/corvallis3d/go-shippo/models"
)

// CreateRefund creates a new refund object.
func (c *Client) CreateRefund(input *models.RefundInput) (*models.Refund, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Refund{}
	err := c.do(http.MethodPost, "/refunds/", input, output)
	return output, err
}

// RetrieveRefund retrieves an existing refund by object id.
func (c *Client) RetrieveRefund(objectID string) (*models.Refund, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Refund{}
	err := c.do(http.MethodGet, "/refunds/"+objectID, nil, output)
	return output, err
}

// ListAllRefunds list all refund objects.
func (c *Client) ListAllRefunds() ([]*models.Refund, error) {
	list := []*models.Refund{}
	err := c.doList(http.MethodGet, "/refunds/", nil, func(v json.RawMessage) error {
		item := &models.Refund{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
