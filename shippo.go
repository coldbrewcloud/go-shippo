package shippo

import "github.com/coldbrewcloud/go-shippo/client"

func NewClient(privateToken string) *client.Client {
	return client.NewClient(privateToken)
}
