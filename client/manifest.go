package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/d5/go-shippo/models"
)

// CreateManifest creates a new manifest object.
func (c *Client) CreateManifest(input *models.ManifestInput) (*models.ManifestOutput, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.ManifestOutput{}
	err := c.do(http.MethodPost, "/manifests/", input, output)
	return output, err
}

// RetrieveManifest retrieves an existing manifest by object id.
func (c *Client) RetrieveManifest(objectID string) (*models.ManifestOutput, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.ManifestOutput{}
	err := c.do(http.MethodGet, "/manifests/"+objectID, nil, output)
	return output, err
}

// ListAllManifests lists all manifest objects.
func (c *Client) ListAllManifests() ([]*models.ManifestOutput, error) {
	list := []*models.ManifestOutput{}
	err := c.doList(http.MethodGet, "/manifests/", nil, func(v json.RawMessage) error {
		item := &models.ManifestOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
