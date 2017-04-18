package shippo

import "github.com/yuderekyu/go-shippo/client"

func NewClient(privateToken string) *client.Client {
	return client.NewClient(privateToken)
}
