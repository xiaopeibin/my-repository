package service

import "my_go_project/service/example"

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
