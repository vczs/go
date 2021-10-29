package main

import "fmt"

//创建一个接口
type interface_5 interface {
	show()
}
//创建一个结构体并实现interface_5接口
type struct_1 struct {
	name string
}
//创建一个绑定struct_1结构体指针实现interface_5接口的方法
func (struct_1 *struct_1) show(){
	fmt.Println("struct_1.show():",struct_1.name)
}
func main(){
	var i5 interface_5
	s1 := struct_1{"hello golang!"}
	i5 = &s1 //当实现接口的结构体方法绑定的是结构体指针时，在给接口赋值结构体时要记得取地址
	i5.show()
}
