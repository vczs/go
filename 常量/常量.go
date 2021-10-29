package main

import "fmt"

/**
常量：
常量使用const修饰
常量在定义时，必须初始化，且不能修改
常量只能修饰bool int系列 float系列 string类型
go中没有常量必须大写的规定
常量通过首字母大小写控制访问范围
语法：const variable [类型] = value    //[类型]可写可不写，因为go有类型推导

举例：
const name = "tom"  //true
const num float64 = 0.8   //true
const a int  //err 没有初始化
const b = 4/2  //true
const c = num/2  //err num不是常量，编译器不能识别
const d = test()  //err test()不是常量，编译器不能识别
 */
func main(){

	//常量简洁写法
	const(
		a = 2
		b = 3
		c    //未初始化的常量默认值为上一个常量的值，即c的值为b的值2
	)
	fmt.Printf("a=%v  b=%v  c=%v",a,b,c)

	//常量专业写法
	const(
		d = iota
		e
		f
		g
		//逻辑：给d赋值为0，e在d的基础上+1，f在e的基础上+1，g在f的基础上+1，以此类推
	)
	fmt.Printf("   d=%v  e=%v  f=%v  g=%v",d,e,f,g)

}
