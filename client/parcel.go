package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/d5/go-shippo/models"
)

// CreateParcel creates a new parcel object.
func (c *Client) CreateParcel(input *models.ParcelInput) (*models.ParcelOutput, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.ParcelOutput{}
	err := c.do(http.MethodPost, "/parcels/", input, output)
	return output, err
}

// RetrieveParcel retrieves an existing parcel by object id.
func (c *Client) RetrieveParcel(objectID string) (*models.ParcelOutput, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.ParcelOutput{}
	err := c.do(http.MethodGet, "/parcels/"+objectID, nil, output)
	return output, err
}

// ListAllParcels lists all parcel objects.
func (c *Client) ListAllParcels() ([]*models.ParcelOutput, error) {
	list := []*models.ParcelOutput{}
	err := c.doList(http.MethodGet, "/parcels/", nil, func(v json.RawMessage) error {
		item := &models.ParcelOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
