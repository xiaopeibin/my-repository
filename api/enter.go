package api

import (
	"my_go_project/api/demo"
	"my_go_project/api/example"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	DemoGroup       demo.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
