package main

import (
	"my_go_project/core"
	"my_go_project/global"
	"my_go_project/initialize"
)

func main() {
	// 读取环境变量
	core.Viper()
	// 初始化日志
	global.GVA_LOG = core.Zap()
	// 初始化数据库连接
	global.GVA_DB = initialize.Gorm()
	// 初始化路由
	core.RunWindowsServer()
}
