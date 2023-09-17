package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/corvallis3d/go-shippo"
	"github.com/corvallis3d/go-shippo/client"
	"github.com/corvallis3d/go-shippo/models"
)

var (
	privateToken     = os.Getenv("PRIVATE_TOKEN")
	upsUserName      = os.Getenv("UPS_USERNAME")
	upsPassword      = os.Getenv("UPS_PASSWORD")
	upsAccountNumber = os.Getenv("UPS_ACCOUNT_NUMBER")
)

func main() {
	// create a Shippo Client instance
	c := shippo.NewClient(privateToken)

	// create or update carrier account
	carrierAccountObjectID := prepareCarrierAccount(c)

	// create shipment using the carrier account
	shipment := createShipmentUsingCarrierAccount(c, carrierAccountObjectID)

	// purchase shipping label
	purchaseShippingLabel(c, shipment)
}

func prepareCarrierAccount(c *client.Client) string {
	// list all registered carrier account
	allCarrierAccounts, err := c.ListAllCarrierAccounts()
	if err != nil {
		panic(err)
	}

	// create UPS carrier account if not added
	carrierAccountObjectID := ""
	for _, ca := range allCarrierAccounts {
		if ca.Carrier == models.CarrierUPS && ca.AccountID == upsUserName {
			carrierAccountObjectID = ca.ObjectID
			break
		}
	}
	if carrierAccountObjectID == "" {
		// create a new carrier account
		input := &models.CarrierAccountInput{
			Carrier:   models.CarrierUPS,
			AccountID: upsUserName,
			Parameters: map[string]interface{}{
				"password":       upsPassword,
				"account_number": upsAccountNumber,
			},
			Active: true,
		}
		carrierAccount, err := c.CreateCarrierAccount(input)
		if err != nil {
			panic(err)
		}

		carrierAccountObjectID = carrierAccount.ObjectID

		fmt.Printf("Carrier account registered: %s\n", dump(carrierAccount))
	} else {
		// update existing carrier account
		input := &models.CarrierAccountInput{
			Carrier:   models.CarrierUPS,
			AccountID: upsUserName,
			Parameters: map[string]interface{}{
				"password":       upsPassword,
				"account_number": upsAccountNumber,
			},
			Active: true,
		}
		carrierAccount, err := c.UpdateCarrierAccount(carrierAccountObjectID, input)
		if err != nil {
			panic(err)
		}

		carrierAccountObjectID = carrierAccount.ObjectID

		fmt.Printf("Carrier account updated: %s\n", dump(carrierAccount))
	}

	return carrierAccountObjectID
}

func createShipmentUsingCarrierAccount(c *client.Client, carrierAccountObjectID string) *models.Shipment {
	// create a sending address
	addressFromInput := &models.AddressInput{
		Name:    "Mr. Hippo",
		Street1: "215 Clayton St.",
		City:    "San Francisco",
		State:   "CA",
		Zip:     "94117",
		Country: "US",
		Phone:   "+1 555 341 9393",
		Email:   "support@goshippo.com",
	}
	addressFrom, err := c.CreateAddress(addressFromInput)
	if err != nil {
		panic(err)
	}

	// create a receiving address
	addressToInput := &models.AddressInput{
		Name:    "Mrs. Hippo",
		Street1: "965 Mission St.",
		City:    "San Francisco",
		State:   "CA",
		Zip:     "94105",
		Country: "US",
		Phone:   "+1 555 341 9393",
		Email:   "support@goshippo.com",
	}
	addressTo, err := c.CreateAddress(addressToInput)
	if err != nil {
		panic(err)
	}

	// create a parcel
	parcelInput := &models.ParcelInput{
		Length:       "5",
		Width:        "5",
		Height:       "5",
		DistanceUnit: models.DistanceUnitInch,
		Weight:       "2",
		MassUnit:     models.MassUnitPound,
	}
	parcel, err := c.CreateParcel(parcelInput)
	if err != nil {
		panic(err)
	}

	// create a shipment
	shipmentInput := &models.ShipmentInput{
		AddressFrom:     addressFrom.ObjectID,
		AddressTo:       addressTo.ObjectID,
		Parcels:         []string{parcel.ObjectID},
		CarrierAccounts: []string{carrierAccountObjectID},
		Async:           false,
	}
	shipment, err := c.CreateShipment(shipmentInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Shipment:\n%s\n", dump(shipment))

	return shipment
}

func purchaseShippingLabel(c *client.Client, shipment *models.Shipment) {
	transactionInput := &models.TransactionInput{
		Rate:          shipment.Rates[0].ObjectID,
		LabelFileType: models.LabelFileTypePDF,
		Async:         false,
	}
	transaction, err := c.PurchaseShippingLabel(transactionInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transaction:\n%s\n", dump(transaction))
}

func dump(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(data)
}
