package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/corvallis3d/go-shippo/models"
)

// RetrieveBatch retrieves an existing batch. BatchShipments are displayed 100 at a time.
// You can iterate through each "page" by specifying non-zero value to page parameter.
// You can also filter based on BatchShipment status using objectResultsFilter parameter
func (c *Client) RetrieveBatch(objectID string, page uint, objectResultsFilter string) (*models.Batch, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

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

	output := &models.Batch{}
	err := c.do(http.MethodGet, url, nil, output)
	return output, err
}

// AddBatchShipmentsToBatch adds batch shipment(s) to an existing Batch.
func (c *Client) AddBatchShipmentsToBatch(objectID string, batchShipments []*models.BatchShipmentInput) (*models.Batch, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}
	if batchShipments == nil {
		return nil, errors.New("nil batch shipments")
	}

	output := &models.Batch{}
	err := c.do(http.MethodPost, "/batches/"+objectID+"/add_shipments", &batchShipments, output)
	return output, err
}

// RemoveBatchShipmentsFromBatch removes batch shipment(s) from an existing Batch.
func (c *Client) RemoveBatchShipmentsFromBatch(objectID string, batchShipmentIDs []string) (*models.Batch, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}
	if batchShipmentIDs == nil {
		return nil, errors.New("nil batch shipment IDs")
	}

	output := &models.Batch{}
	err := c.do(http.MethodPost, "/batches/"+objectID+"/remove_shipments", &batchShipmentIDs, output)
	return output, err
}

// PurchaseBatch purchases an existing batch with an ObjectStatus of "VALID".
// Once you send invoke this function, the batch ObjectStatus will be change to PURCHASING.
// When all the shipments are purchased, the ObjectStatus will change to PURCHASED
// and you will receive a batch_purchased webhook indicating that the batch has been purchased.
func (c *Client) PurchaseBatch(objectID string) (*models.Batch, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Batch{}
	err := c.do(http.MethodPost, "/batches/"+objectID+"/purchase", nil, output)
	return output, err
}
