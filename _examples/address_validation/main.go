package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/corvallis3d/go-shippo"
	"github.com/corvallis3d/go-shippo/models"
)

var (
	privateToken = os.Getenv("PRIVATE_TOKEN")
)

func main() {
	// create a Shippo Client instance
	c := shippo.NewClient(privateToken)

	// Address #1: Valid
	address1Input := &models.AddressInput{
		Name:     "Shawn Ippotle",
		Company:  "Shippo",
		Street1:  "215 Clayton St.",
		City:     "San Francisco",
		State:    "CA",
		Zip:      "94117",
		Country:  "US",
		Email:    "shippotle@goshippo.com",
		Validate: true,
	}
	address1, err := c.CreateAddress(address1Input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address #1: %s\n", dump(address1))

	// Address #2: Wrong street name
	address2Input := &models.AddressInput{
		Name:     "Mr. Hippo",
		Street1:  "4610 Clayton St.",
		City:     "San Francisco",
		State:    "CA",
		Zip:      "94107",
		Country:  "US",
		Email:    "mrhippo@goshippo.com",
		Validate: true,
	}
	address2, err := c.CreateAddress(address2Input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address #2: %s\n", dump(address2))
}

func dump(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(data)
}
