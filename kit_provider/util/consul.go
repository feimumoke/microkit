package util

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/vic/pkg/uid"
	"log"
)

var ConsulClient *api.Client
var ServiceID string
var ServiceName string
var ServicePort int

func init() {
	config := api.DefaultConfig()
	config.Address = "192.168.100.26:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
	ServiceID = "userservice" + uid.New().String()
}

func SetUserServiceNameAndPort(name string, port int) {
	ServiceName = name
	ServicePort = port
}
func RegService() {

	reg := api.AgentServiceRegistration{}
	reg.ID = ServiceID
	reg.Name = ServiceName
	reg.Address = "192.168.100.26"
	reg.Port = ServicePort
	reg.Tags = []string{"primary"}
	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = fmt.Sprintf("http://%s:%d/health",reg.Address,ServicePort)
	reg.Check = &check
	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func Unregservice() {
	ConsulClient.Agent().ServiceDeregister("userservice")
}
