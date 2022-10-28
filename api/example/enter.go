package example

import "my_go_project/service"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService =service.ServiceGroupApp.ExampleServiceGroup.CustomerService
)
