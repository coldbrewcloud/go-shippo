package client

import (
	"errors"
	"net/http"
	"time"

	"github.com/corvallis3d/go-shippo/models"
)

func (c *Client) CreateOrder(input *models.OrderInput) (*models.OrderResponse, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	if input.PlacedAt.IsZero() {
		input.PlacedAt = time.Now()
	}

	output := &models.OrderResponse{}
	err := c.do(http.MethodPost, "/orders/", input, output)
	return output, err
}
