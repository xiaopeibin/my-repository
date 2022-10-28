package demo

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my_go_project/global"
	"my_go_project/model/common/response"
	"my_go_project/model/demo"
)

type SelectApi struct {
}

// SingleSelect 取单条数据
func (*SelectApi) SingleSelect(c *gin.Context) {
	p := demo.Person{}
	global.GVA_DB.First(&p) // 按主键升序排序取第一条
	fmt.Println(p.ToString())

	global.GVA_DB.Last(&p) // 按主键升序排序取最后一条
	fmt.Println(p.ToString())

	global.GVA_DB.Take(&p) // 随机取一条
	fmt.Println(p.ToString())

	result := global.GVA_DB.First(&p, 10) // 根据主键检索(只支持整形数值,防止sql注入)
	fmt.Println(p.ToString())
	if recNotFound(result.Error) {
		response.FailWithMessage("没找到数据!", c)
	}

	resultMap := map[string]interface{}{}
	global.GVA_DB.Model(&demo.Person{}).First(resultMap) // 使用map接收结果
	fmt.Printf("map接收结果:%v", resultMap)

	// 报错,因为不知道主键是那个
	// resultMap1 := map[string]interface{}{}
	// global.GVA_DB.Table("t_person").First(resultMap1)

	// 可以
	resultMap2 := map[string]interface{}{}
	global.GVA_DB.Table("t_person").Take(resultMap2) // 使用table指定查哪个表

	response.OkWithData(p, c)
}

// MultipleSelect 查询多条数据
func (*SelectApi) MultipleSelect(c *gin.Context) {
	var pList []demo.Person
	result := global.GVA_DB.Find(&pList) // 查找多条记录
	if recNotFound(result.Error) {
		response.FailWithMessage("数据未找到!", c)
		return
	}
	fmt.Println("结果行数是:", result.RowsAffected)
	response.OkWithData(pList, c)
}

// WhereSelect 条件查询
func (*SelectApi) WhereSelect(c *gin.Context) {
	p1 := demo.Person{}
	global.GVA_DB.Where("age = ?", "22").First(&p1) // 单条条件查询
	fmt.Println(p1.ToString())

	var pList1 []demo.Person
	global.GVA_DB.Where("name <> ?", "肖培斌").Find(&pList1) // 多条条件查询
	printPersonSlice(pList1)

	var pList2 []demo.Person
	global.GVA_DB.Where("age in ?", []string{"18", "23"}).Find(&pList2) // in
	printPersonSlice(pList2)

	var pList3 []demo.Person
	global.GVA_DB.Where("name like ?", "肖培斌%").Find(&pList3) // like
	printPersonSlice(pList3)

	var pList4 []demo.Person
	global.GVA_DB.Where("name like ? and age = ?", "肖培斌%", "18").Find(&pList4) // and
	printPersonSlice(pList4)

	var pList5 []demo.Person
	global.GVA_DB.Where("age between ? and ?", "20", "30").Find(pList5) // between
	printPersonSlice(pList5)

	response.OkWithData(pList1, c)
}

func (*SelectApi) StructWhereSelect(c *gin.Context) {
	p1 := demo.Person{}
	global.GVA_DB.Where(&demo.Person{ // 结构体作为where条件
		Name: "肖培斌",
		Age:  18,
	}).First(&p1)
	fmt.Println(p1.ToString())

	var pList1 []demo.Person
	global.GVA_DB.Where(map[string]interface{}{ // 声明map作为where条件
		"name": "肖培斌1",
		"age":  30,
	}).Find(&pList1)
	printPersonSlice(pList1)
	// 注:结构体的零值(0,'',false)不参与查询,如想参与查询需要使用map作为参数

	var pList2 []demo.Person
	global.GVA_DB.Where([]int64{3, 4, 5}).Find(&pList2) // 会根据主键 in 来查询
	printPersonSlice(pList2)

	// not标签和where类似
	// https://gorm.cn/zh_CN/docs/query.html#Not-%E6%9D%A1%E4%BB%B6
	response.OkWithData(p1, c)
}

// InnerSelect 内联条件查询
func (*SelectApi) InnerSelect(c *gin.Context) {
	p1 := demo.Person{}
	global.GVA_DB.First(&p1, "id=?", "10") // 单条
	fmt.Println(p1)

	var pList1 []demo.Person
	global.GVA_DB.Find(&pList1, "name like ? and age in ?", "肖%", []int{18, 30}) // 多条查询
	printPersonSlice(pList1)

	var pList2 []demo.Person
	global.GVA_DB.Find(&pList2, demo.Person{Name: "本万"}) // 结构体作为参数
	printPersonSlice(pList2)

	var pList3 []demo.Person
	global.GVA_DB.Find(&pList3, map[string]interface{}{
		"age": 20,
	})
	printPersonSlice(pList3)

	response.Ok(c)
}

// OrSelect 使用or方法
func (*SelectApi) OrSelect(c *gin.Context) {
	var pList []demo.Person
	global.GVA_DB.Where("name = ?", "本万").Or("age = ?", 23).Find(&pList) // or条件查询
	printPersonSlice(pList)

	// or中也可以使用结构体和map作为参数
	// https://gorm.cn/zh_CN/docs/query.html#Or-%E6%9D%A1%E4%BB%B6
	response.OkWithData(pList, c)
}

// SelectField 选择指定的字段
func (*SelectApi) SelectField(c *gin.Context) {
	var persons1 []demo.Person
	global.GVA_DB.Select("name", "age").Find(&persons1) // 只查name和age
	printPersonSlice(persons1)

	var persons2 []demo.Person
	columns := []string{"name", "hobby"} // 要查找的列组成一个切片,作为参数查询
	global.GVA_DB.Select(columns).Find(&persons2)
	printPersonSlice(persons2)

	response.OkWithData(persons2, c)
}

// OrderSelect 排序查询
func (*SelectApi) OrderSelect(c *gin.Context) {
	var pList1 []demo.Person
	global.GVA_DB.Order("age desc,name").Find(&pList1) // 写在一起
	printPersonSlice(pList1)

	var pList2 []demo.Person
	global.GVA_DB.Order("age desc").Order("name").Find(&pList2) // 级联方式
	printPersonSlice(pList2)
	response.Ok(c)
}

// PageSelect limit,offset用法
func (*SelectApi) PageSelect(c *gin.Context) {
	var pList1 []demo.Person
	var pList2 []demo.Person
	global.GVA_DB.Limit(3).Offset(2).Find(&pList1).Limit(-1).Offset(-1).Find(&pList2)
	printPersonSlice(pList1)
	printPersonSlice(pList2)
	response.Ok(c)
}

func recNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func printPersonSlice(pList []demo.Person) {
	fmt.Println("===打印开始===")
	for _, person := range pList {
		fmt.Println(person)
	}
	fmt.Println("===打印结束===")
}
