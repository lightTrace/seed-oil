package util

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
)

var consulClient *api.Client

var serviceId string
var serviceName string
var ServicePort int

func init() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	consulClient = client
	serviceId = "uservice-" + uuid.New().String()
}

func SetServiceInfo(name string, port int) {
	serviceName = name
	ServicePort = port
}

func RegisterService() {

	reg := api.AgentServiceRegistration{}
	reg.ID = serviceId
	reg.Name = serviceName
	reg.Address = "192.168.21.235"
	reg.Port = ServicePort
	reg.Tags = []string{"primary"}

	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = fmt.Sprintf("http://%s:%d/health", reg.Address, reg.Port)

	reg.Check = &check

	err := consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func UnRegisterService() {
	consulClient.Agent().ServiceDeregister(serviceId)
}
