package main

/**
反射介绍:
1.反射在运行时可以动态获取变量的各种信息，比如变量的类型(type)，类别(kind)。
2.如果是结构体变量，还可以获取结构体本身的信息(包括结构体字段、方法)。
3.通过反射，可以修改变量的值，可以调用其绑定的方法。
4.使用反射，需要import "reflect"


反射应用场景:
1.不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口。
2.对结构体序列化时，如果结构体有指定Tag，也会使用到反射生成对应的字符串。


反射的函数和概念：
1.reflect.TypeOf(变量名)，返回reflect.Type类型。
2.reflect.ValueOf(变量名)，返回reflect.Value类型；通过reflect.Value可以获取变量的更多信息。
3.变量、interface{}和reflect.Value是可以相互转换的。
举例：
var stu Student   stu这是一个Student结构体类型的变量
stu在作为参数传给函数时，函数用interface{}数据类型来接收
stu转reflect.Type：  rVal := reflect.ValueOf(stu)
reflect.Type转interface{}：   iVal := rVal.interface()
interface{}转原来变量，使用类型断言：    v := iVal.(Student)


反射注意事项和细节：
1.reflect.Type().Kind()和reflect.Value().Kind()都可以获取变量的类别，返回的是一个常量。
2.Type是类型，Kind是类别；两者可能相同也可能不同。
	var num int = 10 num的Type是int，Kind也是int
	var stu Student  stu的Type是包名.Student，Kind是struct
3.使用反射获取变量的值，要类型匹配；比如修改int类型的变量就应该使用reflect.Value(变量).Int()
4.通过反射修改变量的值，应通过变量的指针完成修改，这样才能改变传入变量的值；使用reflect.Value().Elem()获取传入指针变量的值。
*/