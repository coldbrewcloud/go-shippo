package client

import (
	"encoding/json"
	"net/http"

	"github.com/d5/go-shippo/models"
)

func (c *Client) CreateCustomsItem(input *models.CustomsItemInput) (*models.CustomsItemOutput, error) {
	output := &models.CustomsItemOutput{}
	err := c.do(http.MethodPost, "/customs/items/", input, output)
	return output, err
}

func (c *Client) RetrieveCustomsItem(objectID string) (*models.CustomsItemOutput, error) {
	output := &models.CustomsItemOutput{}
	err := c.do(http.MethodGet, "/customs/items/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllCustomsItems() ([]*models.CustomsItemOutput, error) {
	list := []*models.CustomsItemOutput{}
	err := c.doList(http.MethodGet, "/customs/items/", nil, func(v json.RawMessage) error {
		item := &models.CustomsItemOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}

func (c *Client) CreateCustomsDeclaration(input *models.CustomsDeclarationInput) (*models.CustomsDeclarationOutput, error) {
	output := &models.CustomsDeclarationOutput{}
	err := c.do(http.MethodPost, "/customs/declarations/", input, output)
	return output, err
}

func (c *Client) RetrieveCustomsDeclaration(objectID string) (*models.CustomsDeclarationOutput, error) {
	output := &models.CustomsDeclarationOutput{}
	err := c.do(http.MethodGet, "/customs/declarations/"+objectID, nil, output)
	return output, err
}

func (c *Client) ListAllCustomsDeclaration() ([]*models.CustomsDeclarationOutput, error) {
	list := []*models.CustomsDeclarationOutput{}
	err := c.doList(http.MethodGet, "/customs/declarations/", nil, func(v json.RawMessage) error {
		item := &models.CustomsDeclarationOutput{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
