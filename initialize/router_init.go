package initialize

import (
	"github.com/gin-gonic/gin"
	"my_go_project/middleware"
	"my_go_project/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.GinRecovery())
	exampleRouter := router.RouterGroupApp.Example

	publicGroup := engine.Group("")
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		exampleRouter.InitCustomerRouter(publicGroup)
	}
	return engine
}
