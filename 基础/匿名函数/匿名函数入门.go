package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
}

/**
直接在末尾调用
*/
func case_1() {
	var className int = func(number int) int { //定义一个变量给其赋值一个函数 即创建了一个匿名函数
		number = 100
		return number
	}(20) //在匿名函数后直接传入参数调用
	result := className //className也可以赋值给其他变量
	fmt.Println(result)
}

/**
调用赋值给的变量为函数名
*/
func case_2() {
	className := func(number_1 int) (int, int) {
		number_1 = 100
		number_2 := 101
		return number_1, number_2
	}
	result_1, result_2 := className(99) //匿名函数有两个返回值则需要两个变量来接收
	fmt.Println(result_1, result_2)
}

/**
定义全局变量 函数名为定义的全局变量
*/
var (
	className = func(number_1 int, number_2 int) int { //定义一个全局变量
		return number_1 + number_2
	}
)

func case_3() {
	result := className(10, 20) //全局变量名为函数名
	fmt.Println(result)
}
