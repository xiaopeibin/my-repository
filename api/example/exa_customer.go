package example

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"my_go_project/model/common/response"
	"my_go_project/model/example"
)

type CustomerApi struct {
}

func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer// 声明vo
	err := c.ShouldBind(&customer)// 绑定vo
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = validator.New().Struct(customer)// 创建validator,校验vo
	if err != nil {
		errorss := err.(validator.ValidationErrors)
		response.FailWithMessage(errorss.Error(), c)
		return
	}
	err = customerService.InsertExaCustomer(customer)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
func (e *CustomerApi) UpdateExaCustomer(c *gin.Context)  {}
func (e *CustomerApi) DeleteExaCustomer(c *gin.Context)  {}
func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {}
func (e *CustomerApi) CustomerDetail(c *gin.Context) {
	identityNumber := c.Query("identityNumber")
	detail, err := customerService.GetExaCustomerDetail(identityNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.OkWithData(gin.H{}, c)
			return
		}
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(detail, c)
}
