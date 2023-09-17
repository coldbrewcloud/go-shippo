package main

import (
	"errors"
	"os"

	"github.com/corvallis3d/go-shippo"
)

func main() {
	privateToken := os.Getenv("PRIVATE_TOKEN")
	if privateToken == "" {
		panic(errors.New("Please set $PRIVATE_TOKEN with your Shippo API private token."))
	}

	// create a Shippo Client instance
	c := shippo.NewClient(privateToken)

	// create shipment
	// res := c.CreateOrder()
	// fmt.Println(res)
}
