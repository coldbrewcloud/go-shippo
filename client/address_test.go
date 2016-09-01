package client_test

import (
	"os"
	"testing"

	"github.com/d5/go-shippo/client"
	"github.com/d5/go-shippo/models"
	"github.com/stretchr/testify/assert"
)

func TestClient_CreateAddress(t *testing.T) {
	c := client.NewClient(os.Getenv("PRIVATE_TOKEN"))

	_, err := c.CreateAddress(nil)
	assert.NotNil(t, err)

	_, err = c.CreateAddress(&models.AddressInput{})
	assert.NotNil(t, err)

	_, err = c.CreateAddress(&models.AddressInput{
		ObjectPurpose: "PURCHASE",
		Name:          "Shawn Ippotle",
		Company:       "Shippo",
		Street1:       "215 Clayton St.",
		City:          "San Francisco",
		State:         "CA",
		Zip:           "94117",
		Country:       "US",
		Phone:         "+1 555 341 9393",
		Email:         "shippotle@goshippo.com",
		IsResidential: true,
		Metadata:      "Customer ID 123456",
	})
	assert.Nil(t, err)
}
