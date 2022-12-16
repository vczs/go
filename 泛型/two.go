package main

import "fmt"

// 约束类型F的类型可以为int、string、float64
func PrintfMyVar[F int | string | float64](va []F) {
	// 此时形参va可接收以上约束的三种类型
	fmt.Println(va)
}

func main() {
	str := []string{"aa", "bb", "cc"}
	num := []int{33, 44, 55}
	fl := []float64{33.1, 33.2, 33.4}
	PrintfMyVar(str)
	PrintfMyVar(num)
	PrintfMyVar(fl)
}
