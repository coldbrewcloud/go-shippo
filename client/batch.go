package client

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/d5/go-shippo/models"
)

func (c *Client) RetrieveBatch(objectID string, page uint, objectResultsFilter string) (*models.BatchOutput, error) {
	url := "/batches/" + objectID
	qs := []string{}
	if page > 0 {
		qs = append(qs, fmt.Sprintf("page=%d", page))
	}
	if objectResultsFilter != "" {
		qs = append(qs, "object_results="+objectResultsFilter)
	}
	if len(qs) > 0 {
		url += "?" + strings.Join(qs, "&")
	}

	output := &models.BatchOutput{}
	err := c.do(http.MethodGet, url, nil, output)
	return output, err
}

func (c *Client) AddBatchShipmentsToBatch(objectID string, batchShipments []*models.BatchShipmentInput) (*models.BatchOutput, error) {
	output := &models.BatchOutput{}
	err := c.do(http.MethodPost, "/batches/"+objectID+"/add_shipments", &batchShipments, output)
	return output, err
}

func (c *Client) RemoveBatchShipmentsFromBatch(objectID string, batchShipmentIDs []string) (*models.BatchOutput, error) {
	output := &models.BatchOutput{}
	err := c.do(http.MethodPost, "/batches/"+objectID+"/remove_shipments", &batchShipmentIDs, output)
	return output, err
}

func (c *Client) PurchaseBatch(objectID string) (*models.BatchOutput, error) {
	output := &models.BatchOutput{}
	err := c.do(http.MethodPost, "/batches/"+objectID+"/purchase", nil, output)
	return output, err
}
