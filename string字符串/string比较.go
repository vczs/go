package main

import (
	"fmt"
	"strings"
)

func main() {
	case_8()
}

/**
字符串比较
strings.EqualFold() 比较字符串是否相同，不区分大小写，两个参数:两个要比较的字符串
返回值一个：返回比较结果，相同为true，不相同false
比较字符串是否相同并区分大小写用“==”比较，返回值一个：返回比较结果，相同为true，不相同false
*/
func case_8(){
	result_1 := strings.EqualFold("abc","Abc")
	fmt.Printf("“strings.EqualFold”的比较结果：%v    数据类型：%T\n",result_1,result_1)
	result_2 := "abc" == "Abc"
	fmt.Printf("“==”的比较结果：%v    数据类型：%T",result_2,result_2)
}