package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/coldbrewcloud/go-shippo/models"
)

// CreateShipment creates a new shipment object.
// To create a return label, simply pass the ObjectID of the original outbound Transaction to the ReturnOf field.
// If this flag is set, purchasing the generated rates will create a Scan-based return label.
// The advantage of using a Scan-based label is that you will only be charged if the label is used/scanned
// (not when the label is created).
// The Shippo API currently supports Scan-based returns for USPS, Fedex and UPS.
// When the ReturnOf flag is set, Shippo API will automatically swap the address_from and address_to fields for label creation.
// Please check the return service terms and condition for the carrier you intend to use.
func (c *Client) CreateShipment(input *models.ShipmentInput) (*models.Shipment, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	if input.ShipmentDate.IsZero() {
		input.ShipmentDate = time.Now()
	}

	output := &models.Shipment{}
	err := c.do(http.MethodPost, "/shipments/", input, output)
	return output, err
}

// RetrieveShipment retrieves an existing shipment by object id.
func (c *Client) RetrieveShipment(objectID string) (*models.Shipment, error) {
	if objectID == "" {
		return nil, errors.New("Empty object ID")
	}

	output := &models.Shipment{}
	err := c.do(http.MethodGet, "/shipments/"+objectID, nil, output)
	return output, err
}

// ListAllShipments lists all shipment objects.
func (c *Client) ListAllShipments() ([]*models.Shipment, error) {
	list := []*models.Shipment{}
	err := c.doList(http.MethodGet, "/shipments/", nil, func(v json.RawMessage) error {
		item := &models.Shipment{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
