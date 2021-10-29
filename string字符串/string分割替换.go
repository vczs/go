package main

import (
	"fmt"
	"strings"
)

func main() {
	case_14()
	case_15()
}


/**
字符串中的子字符串替换
strings.Replace()函数：字符串中的字符替换  一个返回值：返回值为替换后的新字符串
四个参数：第一个为字符串，第二个为字符串中要替换的原子字符串，第三个为要替换的新子字符串，第四个为要替换的数量，全部替换填-1
*/
func case_14(){
	str := strings.Replace("go go go 语言","go","golang",2)
	fmt.Println(str)
}


/**
字符分割为数组
strings.Split()函数：字符分割为数组  一个返回值：返回值为分割后的数组
两个参数：第一个为字符串，第二个为分隔符
*/
func case_15(){
	array := strings.Split("hello golang!","ll")
	fmt.Printf("替换后：%v   数据类型：%T",array,array)
}