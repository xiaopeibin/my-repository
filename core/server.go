package core

import (
	"fmt"
	"my_go_project/global"
	"my_go_project/initialize"
)

func RunWindowsServer() {
	// 初始化路由
	engine := initialize.Routers()
	serverPort := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Port)
	engine.Run(serverPort)
}
