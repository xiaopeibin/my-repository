package demo

import (
	"fmt"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name    string   `json:"name" form:"name"`
	Age     int      `json:"age" form:"age"`
	Hobby   []string `json:"hobby" form:"hobby"`
	Address string   `json:"address" form:"address"`
}

func (p Person) ToString() string {
	return fmt.Sprintf("我是一个学生!主键:%d,名字:%s,爱好:%v,住址:%s", p.ID, p.Name, p.Hobby, p.Address)
}
