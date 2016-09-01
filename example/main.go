package main

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/d5/go-shippo/client"
	"github.com/d5/go-shippo/models"
)

func main() {
	client := client.NewClient(os.Getenv("PRIVATE_TOKEN"))

	addressInput := &models.AddressInput{
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
	}

	addressOutput, err := client.CreateAddress(addressInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address created: %s\n", dump(addressOutput))

	retrievedAddress, err := client.RetrieveAddress(addressOutput.ObjectID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address retrieved: %s\n", dump(retrievedAddress))

	allAddresses, err := client.ListAllAddresses()
	if err != nil {
		panic(err)
	}
	for i, a := range allAddresses {
		fmt.Printf("List [%d]: %s\n", i, dump(a))
	}
}

func dump(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(data)
}
