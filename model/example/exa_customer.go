package example

import "gorm.io/gorm"

type ExaCustomer struct {
	gorm.Model
	CustomerIdentityNumber string          `json:"customerIdentityNumber" form:"customerIdentityNumber" gorm:"unique_index,not null,comment:身份证号" validate:"required"`
	CustomerName           string          `json:"customerName" form:"customerName" gorm:"comment:客户名" validate:"required"` // 客户名
	CustomerPhoneData      string          `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`         // 客户手机号
	CreditCardList         []ExaCreditCard `json:"creditCardList" form:"creditCardList" validate:"dive"`
	Address                *Address         `json:"address" form:"address" validate:"omitempty" gorm:"comment:住址"`
}
type ExaCreditCard struct {
	gorm.Model
	ExaCustomerId uint
	CardNumber    string `json:"cardNumber" form:"cardNumber" gorm:"comment:卡号" validate:"required"`
	BankName      string `json:"bankName" form:"bankName" gorm:"comment:银行名称"`
	Overage       int32  `json:"overage" form:"overage" gorm:"comment:余额"`
}
type Address struct {
	gorm.Model
	ExaCustomerId uint
	Province string `json:"province" form:"province" validate:"required"`
	City     string `json:"city" form:"city" validate:"required"`
	Street   string `json:"street" form:"street" validate:"required"`
	Detail   string `json:"detail" form:"detail" `
}
