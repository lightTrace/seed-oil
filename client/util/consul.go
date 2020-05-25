package util

import (
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"log"
)

var consulClient *consul.Client

func init() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	apiClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	client := consul.NewClient(apiClient)
	consulClient = &client
}
