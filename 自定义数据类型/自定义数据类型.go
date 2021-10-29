package main

import "fmt"

func main(){
	case_1()
	case_2()
}

func case_1(){
	type customType int //自定义数据类型,关键字type
	var num1 int = 20 //num1的数据类型是int
	var num2 customType //num2的数据类型是customType
	fmt.Printf("num1的值:%d  数据类型:%T\nnum2的值:%d  数据类型:%T\n",num1,num1,num2,num2)
	num2 = customType(num1) //num1 num2是不同的数据类型,int是int,customType是customType,所以赋值时要强制转换,赋值后值变数据类型不变
	fmt.Printf("~num1的值:%d  数据类型:%T\n~num2的值:%d  数据类型:%T",num1,num1,num2,num2)
}


func case_2(){
	a := case_2_1
	b := case_2_2(a,50,40) //a是case_2_1赋值过来的,所以a可以作为参数传给数据类型为customType的c
	fmt.Printf("a的值:%d  数据类型:%T\nb的值:%d  数据类型:%T\n",a,a,b,b)
}
type customType func(a int , b int) int//customType是自定义数据类型,数据类型为func(a int , b int)
func case_2_1(a int , b int) int {
	return a - b
}
func case_2_2(c customType, a int , b int) int {
	fmt.Printf("c的值:%d  数据类型:%T\n",c,c)
	return a + b
}
