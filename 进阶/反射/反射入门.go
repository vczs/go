package main

import (
	"fmt"
	"reflect"
)

func main() {
	case_1()
	case_2()
}

/**
基本数据类型反射操作
*/
func case_1() {
	var num int64 = 58
	case_1_1(num)
}
func case_1_1(model interface{}) {
	//获取model的reflect.Type
	rType := reflect.TypeOf(model)
	//获取model的reflect.Value
	rVal := reflect.ValueOf(model)
	//获取变量的Kind
	tKind := rType.Kind()
	rKind := rVal.Kind()
	fmt.Printf("tKind=%v(%T)  rKind=%v(%T)\n", tKind, tKind, rKind, rKind)
	//获取model的具体值
	modelValue := rVal.Int()
	//将model转为传入函数之前的变量num：先将rVal转为interface{}类型，再通过类型断言转为变量num
	iVal := rVal.Interface() //将model获取的的reflect.Value转为interface{}类型
	num := iVal.(int64)
	fmt.Printf("rType=%v(%T)\nrVal=%v(%T)\nmodelValue=%v(%T)\niVal=%v(%T)\nnum=%v(%T)\n", rType, rType, rVal, rVal, modelValue, modelValue, iVal, iVal, num, num)
}

/**
结构体反射操作
*/
type Student struct {
	Name string
	Age  int
}

func case_2() {
	stu := Student{
		Name: "hello golang!",
		Age:  18,
	}
	case_2_1(stu)
}
func case_2_1(model interface{}) {
	//获取model的reflect.Type
	rType := reflect.TypeOf(model)
	//获取model的reflect.Value
	rVal := reflect.ValueOf(model)
	fmt.Printf("rType=%v(%T)   rVal=%v(%T)\n", rType, rType, rVal, rVal)
	//获取变量的Kind
	tKind := rType.Kind()
	rKind := rVal.Kind()
	fmt.Printf("tKind=%v(%T)  rKind=%v(%T)\n", tKind, tKind, rKind, rKind)
	//获取model的具体值
	modelValue := rVal.String()
	//将model转为传入函数之前的变量num：先将rVal转为interface{}类型，再通过类型断言转为变量num
	iVal := rVal.Interface() //将model获取的的reflect.Value转为interface{}类型
	str := iVal.(Student)    //类型断言
	strName := str.Name      //str经过类型断言后才能访问它的字段Name
	fmt.Printf("modelValue=%v(%T)\niVal=%v(%T)\nnum=%v(%T)\nstrName=%v(%T)\n", modelValue, modelValue, iVal, iVal, str, str, strName, strName)
}
