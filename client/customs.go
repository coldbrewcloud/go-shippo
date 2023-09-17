package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/corvallis3d/go-shippo/models"
)

// CreateCustomsItem creates a new customs item object.
func (c *Client) CreateCustomsItem(input *models.CustomsItemInput) (*models.CustomsItem, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.CustomsItem{}
	err := c.do(http.MethodPost, "/customs/items/", input, output)
	return output, err
}

// RetrieveCustomsItem retrieves an existing customs item by object id.
func (c *Client) RetrieveCustomsItem(objectID string) (*models.CustomsItem, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.CustomsItem{}
	err := c.do(http.MethodGet, "/customs/items/"+objectID, nil, output)
	return output, err
}

// ListAllCustomsItems lists all customs item objects.
func (c *Client) ListAllCustomsItems() ([]*models.CustomsItem, error) {
	list := []*models.CustomsItem{}
	err := c.doList(http.MethodGet, "/customs/items/", nil, func(v json.RawMessage) error {
		item := &models.CustomsItem{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}

// CreateCustomsDeclaration creates a new customs declaration object.
func (c *Client) CreateCustomsDeclaration(input *models.CustomsDeclarationInput) (*models.CustomsDeclaration, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	switch input.Items.(type) {
	case []string, []*models.CustomsItemInput, nil:
	default:
		return nil, errors.New("invalid items data type")
	}

	output := &models.CustomsDeclaration{}
	err := c.do(http.MethodPost, "/customs/declarations/", input, output)
	return output, err
}

// RetrieveCustomsDeclaration retrieves an existing customs declaration by object id.
func (c *Client) RetrieveCustomsDeclaration(objectID string) (*models.CustomsDeclaration, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.CustomsDeclaration{}
	err := c.do(http.MethodGet, "/customs/declarations/"+objectID, nil, output)
	return output, err
}

// ListAllCustomsDeclaration lists all customs declaration objects.
func (c *Client) ListAllCustomsDeclaration() ([]*models.CustomsDeclaration, error) {
	list := []*models.CustomsDeclaration{}
	err := c.doList(http.MethodGet, "/customs/declarations/", nil, func(v json.RawMessage) error {
		item := &models.CustomsDeclaration{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
