package main

import "fmt"

/**
打印九九乘法表
*/
func main() {
	for i := 1 ; i <= 9 ; i++ {
		for j := 1 ; j <= i ; j++ {
			n := i*j
			fmt.Printf("%v x %v = %v   ",j,i,n) //%v为占位符
			//乘积为个位数时加个空格，表格更加整齐，无其他意义
			if n < 10 {
				fmt.Print(" ")
			}
		}
		fmt.Println()//输出完换行
	}
}