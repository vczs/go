package main

import (
	"fmt"
	"reflect"
)

/**
使用反射遍历结构体的字段，调用结构体的方法，并获取结构体标签的值。
*/

//创建一个结构体
type People struct {
	Name    string `js:"people_name"`
	Age     int    `js:"people_int"`
	Address string `js:"people_address"`
}

//修改结构体各字段的值
func (people People) Set(name string, age int, address string) People {
	people.Name = name
	people.Age = age
	people.Address = address
	return people
}

//打印结构体
func (people People) Show() {
	fmt.Printf("people:%v\n", people)
}

func peopleStruct(people interface{}) {
	//获取结构体的Type
	rTyp := reflect.TypeOf(people)
	//获取结构体的Value
	rVal := reflect.ValueOf(people)

	//获取结构体的Kind
	vKin := rVal.Kind()
	//判断传入的变量是不是结构体，不是就退出。
	if vKin != reflect.Struct {
		fmt.Println("该变量不是结构体！")
		return
	}

	//获取结构体的字段数量
	numFieldStruct := rVal.NumField()
	fmt.Printf("该结构体变量的字段有 %v 个\n", numFieldStruct)

	//遍历结构体的所有字段和字段的标签
	for i := 0; i < numFieldStruct; i++ {
		//获取结构体第i个字段并打印
		field := rVal.Field(i) //返回结构体第i个字段的封装的Value
		fmt.Printf("该结构体变量第 %v 个字段值：%v\n", i+1, field)
		//获取结构体第i个字段的标签并打印 没有就不打印
		tagField := rTyp.Field(i).Tag.Get("js") //这里必须用rTyp.Field(i)获取tag，因为rTyp.Field(i)才会返回结构体的第i个字段
		if tagField != "" {
			fmt.Printf("该结构体变量第 %v 个字段的标签：%v\n", i+1, tagField)
		}
	}

	//获取结构体的方法数量并打印（只能获取首字母大写的方法）
	numMethodStruct := rVal.NumMethod()
	fmt.Printf("该结构体变量的方法有 %v 个\n", numMethodStruct)
	//调用它的第二个方法 方法排名顺序是按照方法名在ASCII码表中的顺序进行顺序
	rVal.Method(1).Call(nil)
	//调用它的Set方法,传入参数是一个切片,先创建一个切片
	//创建一个切片
	var slice []reflect.Value
	slice = append(slice, reflect.ValueOf("hello golang!"))
	slice = append(slice, reflect.ValueOf(22))
	slice = append(slice, reflect.ValueOf("US"))
	peopleSet := rVal.MethodByName("Set").Call(slice) //调用它的Set方法,传入参数是一个切片
	fmt.Printf("peopleSet:%v\n", peopleSet[0])
}

func main() {
	people := People{
		Name:    "golang",
		Age:     21,
		Address: "Google",
	}
	peopleStruct(people)
}
