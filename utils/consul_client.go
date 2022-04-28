package utils

import (
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
)

type consul struct {
	client *api.Client
	prefix string
}

func NewConsul(prefix string) *consul {
	address := "consul.hho-inc.com"
	port := "80"
	conf := api.DefaultConfig()
	conf.Address = address + ":" + port

	client, err := api.NewClient(conf)
	cobra.CheckErr(err)

	return &consul{
		client, prefix,
	}
}

func (c *consul) GetKV(app string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	KVPair, _, err := c.client.KV().Get(c.prefix+"/"+app, &api.QueryOptions{})
	return KVPair.Value, err
}

