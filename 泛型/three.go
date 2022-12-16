package main

import "fmt"

// 泛型方法
type MySlice[F string | float64] []F

func (m MySlice[F]) Add() F {
	var sum F
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	var aa MySlice[string] = []string{"aaa", "bbb", "ccc"}
	fmt.Println(aa.Add())
	var bb MySlice[float64] = []float64{11.1, 11.2, 11.3}
	fmt.Printf("%.2f\n", bb.Add())
}
