package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
}

//创建一个空接口A
type A interface{}
//创建一个结构体B
type B struct {
	name string
}
func case_1(){
	var a A
	var b B = B{"结构体b"}
	a = b //一个空的接口没有定义任何方法，可以说所有的类型都实现了该接口，b实现了a所以b可以赋值给a，空接口可以接收任意类型
	var c B //再定义一个B数据类型的变量c
	//c = a.(B)就是类型断言，逻辑是：判断接口变量a是否指向B数据类型的变量，如果是就将接口变量a转成B数据类型变量并赋值给c，否则报错
	c = a.(B) //c必须是接口a指向的变量b的数据类型B,即变量c必须是()里的数据类型,接口a必须指向()里数据类型的变量
	fmt.Println(c)
}


func case_2(){
	var a interface{} //空接口
	var b int = 12
	a = b //空接口可以接收任意类型
	c := a.(int) //c必须是接口a指向的变量b的数据类型int,即变量c必须是()里的数据类型,接口a必须指向()里数据类型的变量
	fmt.Printf("c的值:%v  数据类型:%T",c,c)
}


/**
带检测的类型断言
给断言语句加个判断，无论结果如何，保障后续代码继续执行
 */
func case_3(){
	var a interface{} //空接口
	var b float64 = 5.8
	a = b //空接口可以接收任意类型
	//c断言后的值  result断言结果：成功为true,失败为false
	if c,result := a.(float64) ; result {
		fmt.Printf("断言成功：c的值:%v  数据类型:%T",c,c)
	} else {
		fmt.Printf("断言失败：c的值:%v  数据类型:%T",c,c)
	}
	fmt.Println("\n程序继续执行。。。")
}