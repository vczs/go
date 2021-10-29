package main

import (
	"fmt"
	"strings"
)

func main() {
	case_16()
	case_17()
}

/**
字符串去前后空格
strings.TrimSpace()函数：字符串去前后空格  一个参数：要处理的原字符串  一个返回值：处理后的新字符串
*/
func case_16(){
	str := " hello golang!  "
	fmt.Println(strings.TrimSpace(str))
}


/**
去除字符串开头或结尾指定子字符串
strings.Trim()函数：去除字符串开头和结尾指定子字符串  两个参数：第一个为原字符串，第二个为字符串中要去的子字符串  一个返回值：处理后的新字符串
strings.TrimLeft()函数：去除字符串开头指定子字符串  两个参数：第一个为原字符串，第二个为字符串中要去的子字符串  一个返回值：处理后的新字符串
strings.TrimRight()函数：去除字符串结尾指定子字符串  两个参数：第一个为原字符串，第二个为字符串中要去的子字符串  一个返回值：处理后的新字符串
*/
func case_17(){
	str := "!hello golang!"
	fmt.Println(strings.Trim(str,"!"))//去除开头和结尾指定字符
	fmt.Println(strings.TrimLeft(str,"!he"))//去除开头指定字符
	fmt.Println(strings.TrimRight(str,"g!"))//去除结尾指定字符
}
