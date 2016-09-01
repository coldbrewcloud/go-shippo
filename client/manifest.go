package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) CreateManifest(input *models.ManifestInput) (*models.ManifestOutput, error) {
	output := &models.ManifestOutput{}
	err := c.do(http.MethodPost, "/manifests/", input, output)
	return output, err
}

func (c *Client) RetrieveManifest(objectID string) (*models.ManifestOutput, error) {
	output := &models.ManifestOutput{}
	err := c.do(http.MethodGet, "/manifests/"+objectID, nil, output)
	return output, err
}

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
