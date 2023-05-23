package main

import (
	"fmt"
	"strings"
)

func main() {
	case_9()
	case_10()
	case_11()
	case_12()
}

/**
字符串是否包含子字符串
strings.Contains()函数  两个参数：第一个参数 字符串，第二个参数 子字符串
函数返回值一个：布尔值，包含就返回true，不包含返回false
*/
func case_9(){
	result := strings.Contains("hello golang!","ll")
	fmt.Printf("结果：%v    数据类型：%T",result,result)
}


/**
统计字符串中子字符串出现的次数
strings.Contains()函数  两个参数：第一个参数 字符串，第二个参数 要统计的子字符串
函数返回值一个：子字符串出现的次数
*/
func case_10(){
	result := strings.Count("hello golang!","l")
	fmt.Printf("结果：%v    数据类型：%T",result,result)
}


/**
字符串在字符串中出现的下标，没有就返回-1
strings.Index()函数计算子字符串在字符串中首次出现的下标  一个返回值：返回值为子字符串出现的下标，没有出现为-1
strings.LastIndex()函数计算子字符串在字符串中最后一次出现的下标  一个返回值：返回值为子字符串出现的下标，没有出现为-1
两个函数都是两个参数：第一个为字符串，第二个为要计算的子字符串
*/
func case_11(){
	//首次出现
	sub_start := strings.Index("hello golang!","ll")
	fmt.Println(sub_start)
	//最后一次出现
	sub_end := strings.LastIndex("hello golang!","an")
	fmt.Println(sub_end)
}




/**
判断字符串是否以指定子字符串开头和结尾
strings.HasPrefix()函数：判断字符串是否以指定子字符串开头  两个参数：第一个为原字符串，第二个为要判断的子字符串  一个返回值：判断结果，布尔值
strings.TrimLeft()函数：判断字符串是否以指定子字符串结尾  两个参数：第一个为原字符串，第二个为要判断的子字符串  一个返回值：判断结果，布尔值
*/
func case_12(){
	str := "hello golang!"
	fmt.Printf("结果：%v   数据类型：%T\n",strings.HasPrefix(str,"he"),strings.HasPrefix(str,"he"))//是否以指定字符开头
	fmt.Printf("结果：%v   数据类型：%T",strings.HasSuffix(str,"png"),strings.HasSuffix(str,"e!"))//是否以指定字符结尾
}
