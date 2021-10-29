package main

import (
	"fmt"
)

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
}

//数组的初始化
func case_1(){
	//一
	var array_a [3]int
	array_a[0] = 7
	array_a[1] = 8
	array_a[2] = 9
	//二
	var array_b [3]int = [3]int{2: 9,1:8,0:7}
	//三
	var array_c = [3]int{2: 9,0:7,8}
	//四
	array_d := [3]int{0: 7,8}
	//五
	var array_e = [...]int{7,2:8,9}
	//六
	var array_f = [...]int{1: 7,3:9}
	fmt.Println("array_a:", array_a,"\narray_b:", array_b,"\narray_c:", array_c,"\narray_d:", array_d,"\narray_e:", array_e,"\narray_f:", array_f)
}


/**
数组的遍历（for循环遍历）
*/
func case_2(){
	var array [26]byte //byte数据类型的数组
	n := 0
	for i := 65 ; i < 91 ; i++ {
		array[n] = byte(i) //i是int数据类型的，array是byte数组，i赋值给数组元素要强转
		fmt.Printf("%c", array[n])
		n++
	}
}


/**
数组的遍历（for range遍历）
 */
func case_3(){
	var array [4]float64
	array[0] = 10.2
	array[1] = 8.8
	array[2] = 11.0
	array[3] = 15.0
	sum := 0.0
	for index,value := range array {
		sum += array[index]
		fmt.Println(index,value)
	}
	average := sum / float64(len(array))  //sum为float64数据类型，len(array)是数组长度是int数据类型，go中不同数据类型运算必须强转
	fmt.Printf("平均数：%.2f\n总和：%.2f",average,sum)  //%.2f：精确到小数点后两位
}


/**
数组的数据类型和数组的len也有关
len为2的数组和len为3的数组，不是同一数据类型
参数是 len为3的数组 的函数，调用时传入参数 len为2的数组 会报错,即使数组元素的数据类型一致
 */
func case_4(){
	array := [3]int{1,2,3}
	case_4_1(array)
}
func case_4_1(array [3]int){
	fmt.Print(array)
}