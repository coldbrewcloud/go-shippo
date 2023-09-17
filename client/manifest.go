package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/corvallis3d/go-shippo/models"
)

// CreateManifest creates a new manifest object.
func (c *Client) CreateManifest(input *models.ManifestInput) (*models.Manifest, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Manifest{}
	err := c.do(http.MethodPost, "/manifests/", input, output)
	return output, err
}

// RetrieveManifest retrieves an existing manifest by object id.
func (c *Client) RetrieveManifest(objectID string) (*models.Manifest, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Manifest{}
	err := c.do(http.MethodGet, "/manifests/"+objectID, nil, output)
	return output, err
}

// ListAllManifests lists all manifest objects.
func (c *Client) ListAllManifests() ([]*models.Manifest, error) {
	list := []*models.Manifest{}
	err := c.doList(http.MethodGet, "/manifests/", nil, func(v json.RawMessage) error {
		item := &models.Manifest{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
