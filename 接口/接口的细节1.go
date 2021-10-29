package main

import "fmt"

func main(){
	case_1()
	case_2()
}

type interface_1 interface {
	show()
}
type interface_2 struct {
	name string
}
func (i2 interface_2)show(){
	fmt.Println("i2.show()：",i2.name)
}
func case_1(){
	var i1 interface_1
	i2 := interface_2{"hello golang"}
	i1 = i2 //一个自定义类型只有实现了某个接口，才能将自身的实例赋值给该接口再被调用
	i1.show()
}


type interface_3 interface {
	show()
}
type interface_4 int //不仅仅是结构体，只要是自定义数据类型，都可以实现接口
func (i4 interface_4)show(){
	fmt.Println("i4.show():i4 =",i4)
}
func case_2(){
	var i3 interface_3
	var i4 interface_4 = 20
	i3 = i4 //一个自定义类型只有实现了某个接口，才能将自身的实例赋值给该接口再被调用
	i3.show()
}