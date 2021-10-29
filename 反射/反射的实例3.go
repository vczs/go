package main

import (
	"fmt"
	"reflect"
)

/**
使用反射创建并操作结构体：
model是User结构体的指针  rType是通过reflect.TypeOf(model)获取的model指针的Type
此时获取的rType是指针变量的Type  要想获取具体值就要通过rType.Elem()获取指向的值
获取到具体值后通过reflect.New(rTypeValue)获取到User结构体的Value，返回的rValue也是个指针类型
再将指针类型的rValue通过rValue.Elem()获取到指向的值rValueValue
通过*model的Value.Elem()即rValueValue就可以对User结构体进行赋值
 */

type User struct{
	Admin string
	Password string
}

func main(){
	var(
		model *User //User指针类型的变量model
		rType reflect.Type
		rValue reflect.Value
	)

	rType = reflect.TypeOf(model) //获取类型为*User的model变量的Type
	fmt.Printf("rType的类别是：%v\n",rType.Kind())
	rTypeValue := rType.Elem() //获取rType指向的值
	fmt.Printf("rType所指向：%v  类型是：%v\n",rTypeValue,rTypeValue.Kind())

	//返回一个Value类型值，它的类别是指针； 它持有一个指向类型为User结构体的新指针
	rValue = reflect.New(rTypeValue) //此时rValue指向一个User类型的空结构体
	fmt.Printf("rValu的类别是：%v\n",rValue.Kind())
	rValueValue := rValue.Elem() //获取rValue指向的值,此时rValueValue是一个空的User结构体变量
	fmt.Printf("rValu所指向：%v  类型是：%v\n",rValueValue,rValueValue.Kind())

	//将model转为传入函数之前的变量*User：先将rValue转为interface{}类型，再通过类型断言转为变量*User
	//类型断言 rValue.Interface() 是否为 *User 类型  断言成功就可以对model进行和User结构体一样的操作
	model = rValue.Interface().(*User) //此时的model是一个指针类型，指向User类型
	fmt.Printf("model的类型是：%v   指向：%v\n",model,*model)

	//给获取的rValue指向的值 User 结构体变量赋值
	//因为是结构体，所以赋值要分别给它的字段进行赋值
	rValueValue.FieldByName("Admin").SetString("golang")
	rValueValue.FieldByName("Password").SetString("hello World！")

	//理解：rValueValue就是结构体指针model的Value，利用反射可以通过结构体指针model的Value.Elem()对model所指向的结构体进行赋值
	//最后打印*model就可以这个赋值后的User结构体
	fmt.Printf("rValueValue = %v    rValueValue.Admin = %v\n",*model,(*model).Admin) //打印(*model)检查是否创建并给其字段赋值成功
}