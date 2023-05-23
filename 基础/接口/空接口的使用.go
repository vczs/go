package main

import "fmt"

//创建一个空接口
type interface_6 interface {

}
//随便创建一个变量
type struct_2 struct {
	name string
}
func main() {
	var i6 interface_6 //一个空的接口没有定义任何方法，可以说所有的类型都实现了该接口。
	var s2 struct_2
	i6 = s2
	fmt.Println(i6)

	//接口也可以作为实例化变量时的数据类型
	var i7 interface{} //interface{}是一个空接口数据类型，{}里可以添加接口要定义的方法
	//结构体也可以作为实例化变量时的数据类型，{}里可以添加结构体的字段（这样创建的结构体只可以添加一个字段）
	var s3 struct{}
	i7 = s3
	fmt.Println(i7)
}
