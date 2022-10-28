package router

import "my_go_project/router/example"

type RouterGroup struct {
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
