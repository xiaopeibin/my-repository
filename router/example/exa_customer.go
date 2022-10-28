package example

import (
	"github.com/gin-gonic/gin"
	"my_go_project/api"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(router *gin.RouterGroup) {
	customerRouter := router.Group("customer")
	exaCustomerApi := api.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // 创建客户
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouter.GET("customerDetail",exaCustomerApi.CustomerDetail)
		customerRouter.GET("customerList", exaCustomerApi.GetExaCustomerList)
	}
}
