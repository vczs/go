package main

import (
	"fmt"
	"reflect"
)

/**
使用反射遍历结构体的字段，并修改字段。
使用reflect.Value().Elem()函数来获取指针变量所指的值
*/

//创建一个结构体
type Monster struct {
	Name    string `js:"monster_name"`
	Age     int    `js:"monster_int"`
	Address string `js:"monster_address"`
}

func monsterStruct(monster interface{}) {
	//获取结构体指针的Type
	rTyp := reflect.TypeOf(monster)
	//获取结构体指针的Value
	rVal := reflect.ValueOf(monster)

	//获取结构体指针的Kind
	vKin := rVal.Kind()
	//如果传入的变量不是结构体指针类型，或传入的变量指向的不是结构体类型，就return函数。
	if vKin != reflect.Ptr || rVal.Elem().Kind() != reflect.Struct {
		fmt.Println("该变量不是结构体指针类型！")
		return
	}

	//获取该结构体指针变量所指的结构体的字段数
	numFieldStruct := rVal.Elem().NumField()
	fmt.Printf("该结构体指针变量所指的结构体的字段有 %v 个\n", numFieldStruct)

	//遍历该结构体指针变量所指的结构体的所有字段和字段的标签
	for i := 0; i < numFieldStruct; i++ {
		//获取第i个字段并打印
		field := rVal.Elem().Field(i) //返回第i个字段的封装的Value
		fieldType := field.Kind()     //获取该字段的类别
		fmt.Printf("该结构体变量第 %v 个字段：%v (%v)\n", i+1, field, fieldType)
		//获取第i个字段的标签并打印 没有就不打印
		tagField := rTyp.Elem().Field(i).Tag.Get("js") //这里必须用rTyp.Field(i)获取tag，因为rTyp.Field(i)才会返回结构体的第i个字段
		if tagField != "" {
			fmt.Printf("该结构体变量第 %v 个字段的标签：%v\n", i+1, tagField)
		}
	}

	//修改该结构体指针变量所指的结构体的第一个字段
	//因为第一个字段数据类型为string，所以使用SetString()函数；如果是int类型，则使用SetInt()进行修改。
	rVal.Elem().Field(0).SetString("hello golang！")
	//修改该结构体指针变量所指的结构体的"Age"字段
	rVal.Elem().FieldByName("Age").SetInt(18)
}

func main() {
	monster := Monster{
		Name:    "golang",
		Age:     21,
		Address: "Google",
	}
	fmt.Println("修改前的monster = ", monster) //打印修改前的结构体
	monsterStruct(&monster)                //传入monster地址
	fmt.Println("修改后的monster = ", monster) //打印修改后的结构体，检查是否修改成功
}
