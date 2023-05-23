package main

import "fmt"

/**
斐波拉契数列
给出要求的斐波拉契数列的个数，控制台打印对应个数的斐波拉契数列
斐波那契数列：1 1 2 3 5 8....构成了一个序列,这个数列的特点：前面相邻两项之和，构成了后一项。
*/
func main(){
	var number int
start:fmt.Println("输入个数：")
	fmt.Scan(&number)
	if fibonacci(number) == nil {
		fmt.Println("输入有误！（大于等于2）")
		goto start
	}else{
		fmt.Print("result：",fibonacci(number))
	}
}
func fibonacci(number int)([]int){
	if number < 2 {
		return nil //如果小于2就返回空
	}
	slice := make([]int, number, number)
	slice[0] = 1
	slice[1] = 1
	for i := 2 ; i < number; i++{
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}