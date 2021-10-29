package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
}

/**
二维数组的创建方式
*/
func case_1() {
	//一：先创建再赋值,为赋值的为默认值0
	var tdArray_1 [3][3]int
	tdArray_1[0][0] = 0
	tdArray_1[0][1] = 1
	tdArray_1[0][2] = 2
	tdArray_1[1][0] = 3
	tdArray_1[2][1] = 4
	//二：创建并赋值
	tdArray_2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("tdArray_1：%v   数据类型：%T\n", tdArray_1, tdArray_1)
	fmt.Printf("tdArray_2：%v   数据类型：%T", tdArray_2, tdArray_2)
}

/**
二维数组双for循环嵌套遍历
*/
func case_2() {
	tdArray := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(tdArray); i++ {
		for j := 0; j < len(tdArray[i]); j++ {
			fmt.Println(tdArray[i][j])
		}
	}
}

/**
二维数组for range遍历
*/
func case_3() {
	tdArray := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for indexOne, valueOne := range tdArray {
		fmt.Println("arrayOne:", indexOne, valueOne)        //indexOne为一维数组下标，valueOne是一个数组，即二维数组
		for indexTwo, valueTwo := range tdArray[indexOne] { //遍历一维数组的值得到二维数组的下标和值
			fmt.Println("arrayTwo:", indexTwo, valueTwo) //indexTwo为二维数组的下标，valueTwo是一个整数，也就是二维数组的值
		}
	}
}
