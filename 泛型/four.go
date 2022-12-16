package main

import "fmt"

// 自定义泛型
type MyFanXing interface {
	int | int8 | ~int64 | float64 | string
}

// F的约束是MyFanXing   MyFanXing是自定义泛型
func GetMaxNum[F MyFanXing](a, b F) F {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n1 int = 188
	var n2 int = 12
	fmt.Println(GetMaxNum(n1, n2))

	// 使用自定义类型时需在自定义泛型中找到其基础类型 并在前面加“~”
	// 下面myInt类型的基础类型是int64  所以要在MyFanXing中的int64前面加“~” 这样传n3,n4才不会报错
	type myInt int64
	var n3 myInt = 15
	var n4 myInt = 111
	fmt.Println(GetMaxNum(n3, n4))
}
