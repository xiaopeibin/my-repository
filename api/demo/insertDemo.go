package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_go_project/global"
	"my_go_project/model/common/response"
	"my_go_project/model/demo"
)

type InsertApi struct {
}

// SimpleInsert 单独插入
func (*InsertApi) SimpleInsert(c *gin.Context) {
	p1 := demo.Person{
		Name: "肖培斌",
		Age:  18,
		Hobby: []string{
			"睡觉", "玩",
		},
		Address: "辽宁省大连市",
	}
	global.GVA_DB.Create(&p1)
	fmt.Println(p1.ToString())

	p2 := demo.Person{
		Name: "肖培斌1",
		Age:  30,
		Hobby: []string{
			"睡觉", "玩",
		},
		Address: "辽宁省大连市",
	}
	global.GVA_DB.Select("Name", "Address").Create(&p2)
	fmt.Println(p2.ToString())

	p3 := demo.Person{
		Name: "肖培斌3",
		Age:  33,
		Hobby: []string{
			"睡觉", "玩",
		},
		Address: "辽宁省大连市3",
	}
	global.GVA_DB.Omit("Hobby").Create(&p3)
	fmt.Println(p3.ToString())
	response.OkWithData(p1, c)
}

// SimpleInsertMap 使用map插入一条数据
func (*InsertApi) SimpleInsertMap(c *gin.Context) {
	insertMap := map[string]interface{}{
		"name": "黄鹤",
		"age":  20,
		"hobby": []string{
			"唱", "跳",
		},
		"address": "辽宁省大连市",
	}
	global.GVA_DB.Table("t_person").Create(&insertMap)
	global.GVA_DB.Model((*demo.Person)(nil)).Create(&insertMap)
	global.GVA_DB.Model(&demo.Person{}).Create(&insertMap)
	response.OkWithData(insertMap, c)
}

func (*InsertApi) MultipleInsert(c *gin.Context) {
	p1 := demo.Person{
		Name: "屈强",
		Age:  23,
		Hobby: []string{
			"vue", "game",
		},
		Address: "辽宁省大连市",
	}
	p2 := demo.Person{
		Name: "郑灿辉",
		Age:  20,
		Hobby: []string{
			"sleep", "game",
		},
		Address: "辽宁省大连市",
	}
	p3 := demo.Person{
		Name: "本万",
		Age:  30,
		Hobby: []string{
			"drive", "eat",
		},
		Address: "辽宁省大连市",
	}
	personSlice := []demo.Person{
		p1, p2, p3,
	}
	global.GVA_DB.Create(&personSlice)
	for _, person := range personSlice {
		fmt.Println(person.ToString())
	}
	response.Ok(c)
}
