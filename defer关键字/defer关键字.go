package main

import "fmt"

func main(){
	fmt.Println(case_1())
}

func case_1() int {
	//defer语句后面的指令先不执行，进入defer栈 遵循先进后出原则
	//等整个函数里的内容执行完毕才开始执行defer栈的内容，执行完defer栈内容才算该函数执行完毕
	defer fmt.Println("11111")
	defer fmt.Println("22222")
	fmt.Println("33333")
	return 00000
}