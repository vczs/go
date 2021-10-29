package main

import "fmt"


func main(){
	case_1()
	case_2()
}


/**
未知函数参数个数但相同数据类型
 */
func case_1(){
	case_1_1(1,5,7,9)
}
func case_1_1(number... int){
	sum := 0
	for i := 0 ; i < len(number) ; i++{
		sum += number[i]
	}
	fmt.Println("参数总和：",sum)
}


/**
自定义函数参数数据类型
 */
func case_2(){
	temp := case_2_1 //函数可以赋值给另一个变量,赋值时可以不传入参数,只需函数名
	result_1,result_2 := temp(11,22)     //函数调用时要传入参数,赋值的变量同样需要
	result_3 := case_2_2(case_2_1,33,44) //函数作为变量时无需传入参数,调用时必须传入参数
	fmt.Print("result_1:",result_1,"    ","result_2:",result_2,"    ","result_3:",result_3)
}
/**
a,b 为两个函数参数
int为返回值数据类型,两个及两个以上用括号括起来，一个不用
 */
func case_2_1(a int , b int) (int , int){
	return a+b,a-b
}
/**
myParameter是传入参数的名称,因为传入的是一个函数所以要为其起名,函数也可以作为数据类型
func(a int , b int) (int , int)为myParameter参数的数据类型,传入时也要保持数据类型对应,函数"类型_1"与之数据类型对应
声明参数myParameter为func(a int , b int) (int , int)数据类型,如同声明参数c为int数据类型
c,d为第二个和第三个参数
"类型_2"函数返回值只有一个，所以不需要用括号括起来
 */
func case_2_2(myParameter func(a int , b int) (int , int), c int , d int) int {
	return c + d
}