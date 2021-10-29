package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
}

/**
结构体的几种创建和赋值方式
使用和赋值结构体前先要创建结构体
*/
type Cat struct { //声明结构体格式：type 结构体名 struct {结构体字段}
	name  string
	age   int
	color string
}

func case_1() {
	var cat_1 Cat
	cat_1.name = "golang"
	cat_2 := Cat{
		name:  "cat_2",
		age:   15,
		color: "read",
	}
	cat_3 := Cat{"go", 18, "blue"}
	var cat_4 *Cat = new(Cat)
	(*cat_4).name = "cat_4"
	(*cat_4).age = 4
	var cat_5 *Cat = &Cat{
		name:  "cat_5",
		age:   5,
		color: "red",
	}
	//结构体也可以作为实例化变量时的数据类型，{}里可以添加结构体的字段（这样创建的结构体只可以添加一个字段）
	var dog struct{ name string }
	dog.name = "dog"
	fmt.Print(cat_1, cat_2, cat_3, *cat_4, *cat_5, dog)
}

/**
结构体各种数据类型字段的赋值
*/
type Person struct {
	name   string
	age    int
	scores [3]int
	slice  []int
	map1   map[string]string
}

func case_2() {
	var person Person
	person.name = "person"
	person.age = 12
	person.scores[0] = 174
	person.scores[1] = 52
	person.slice = make([]int, 2)
	person.slice[1] = 100
	person.map1 = make(map[string]string)
	person.map1["一"] = "111"
	person.map1["二"] = "222"
	fmt.Print(person)
}

/**
结构体在内存中的存储顺序是连续的
*/
type Point struct {
	x, y int
}
type Rect struct {
	up, down Point //Rect成员的数据类型是 Point结构体 数据类型
}

func case_3() {
	line := Rect{Point{1, 2}, Point{5, 6}}
	//按创建和赋值的顺序打印他们的内存地址，可以看到他们的地址在内存中是连续的（计算机内存采用的是16进制）
	fmt.Printf("line.up.x的地址：%p\nline.up.y的地址：%p\nline.down.x的地址：%p\nline.down.y的地址：%p\n", &line.up.x, &line.up.y, &line.down.x, &line.down.y)
	fmt.Printf("line的数据类型：%T\nline.up的数据类型：%T\nline.up.x的数据类型：%T\n", line, line.up, line.up.x)
}

/**
结构体相互转换：两个结构体的字段完全相同的话，这两个结构体就可以相互转换（字段名称 字段数据类型 字段的数量都必须相同）
*/
type A struct {
	name string
	age  int
}
type B struct {
	name string
	age  int
}
type C A //自定义一个数据类型：名称是C，C的数据类型是A
func case_4() {
	a := A{
		name: "a",
		age:  1,
	}
	b := B{"b", 2}
	var c C
	fmt.Println(a, b, c)
	a = A(b) //A和B的字段的数量、名称、数据类型都完全相同，可以转换
	c = C(b) //A和B的字段的数量、名称、数据类型都完全相同，又因为A和C也完全相同，所以C和B完全形同，可以转换
	fmt.Print(a, b, c)
}
