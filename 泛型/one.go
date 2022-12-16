package main

import "fmt"

// 声明泛型变量

func main() {
	// 声明一个F类型切片
	// F约束int、string类型
	type Slice[F int | string] []F
	var a Slice[int] = []int{11, 12, 13}
	var b Slice[string] = []string{"11", "22", "33"}
	fmt.Printf("a=%v 类型:%T\nb=%v 类型:%T\n", a, a, b, b)

	// 声明一个key为F1类型value为F2类型的map
	// F1约束int、string类型  F2约束float64、string类型
	type MyMap[F1 int | string, F2 float64 | string] map[F1]F2
	var c MyMap[int, float64] = map[int]float64{10: 10.06, 11: 12.3}
	fmt.Printf("c=%v 类型:%T\n", c, c)

	// 声明一个ID字段为F类型的结构体
	type MyStruct[F int | string] struct {
		ID   F
		Name string
	}
	d := MyStruct[string]{
		ID:   "001",
		Name: "golang",
	}
	fmt.Printf("d=%v 类型:%T\n", d, d)
}
