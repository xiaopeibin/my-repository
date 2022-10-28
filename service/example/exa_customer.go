package example

import (
	"my_go_project/global"
	"my_go_project/model/example"
)

type CustomerService struct {
}

func (exa *CustomerService) InsertExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *CustomerService) GetExaCustomerDetail(identityNumber string) (detail example.ExaCustomer, err error) {
	customer := example.ExaCustomer{}
	dbRes := global.GVA_DB.Preload("CreditCardList").Preload("Address").First(&customer, "customer_identity_number=?", identityNumber)
	return customer, dbRes.Error
}

//func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
//	err = global.GVA_DB.Delete(&e).Error
//	return err
//}

//func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
//	err = global.GVA_DB.Save(e).Error
//	return err
//}
