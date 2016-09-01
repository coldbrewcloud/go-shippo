package shippo

import "github.com/d5/go-shippo/client"

func NewClient(privateToken string) *client.Client {
	return client.NewClient(privateToken)
}
