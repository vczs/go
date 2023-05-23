package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
	case_5()
	case_6()
}

/**
字符串转整数
strconv.Atoi()函数用来将字符串转为整数，括号里填要转整数的字符串
 */
func case_1(){
	number,err := strconv.Atoi("1234") //number接收转换后的整数，err接收转换失败的原因
	if err != nil{ //err不为空，发生错误，转换失败
		fmt.Println("发生错误：",err) //转换失败时number为默认值0
	}else{ //err为空表示转换无错误，转换成功
		fmt.Println(number)
	}
}


/**
整数转字符串
strconv.Itoa()函数整数转字符串，函数里填要转为字符串的整数
 */
func case_2(){
	str := strconv.Itoa(-123) //函数只有一个返回值，str接收转换后的字符串；整数转字符串必转成功不会发生错误，因此没有错误原因返回值
	fmt.Printf("转换结果：%v    数据类型：%T",str,str)
}


/**
字符串转切片
slice := []byte()函数将字符串转为切片，转换为byte数据类型的切片，括号里写要转为切片的字符串
 */
func case_3(){
	slice := []byte("hello golang!") //slice接收转换后的切片,字符串转切片必转成功不会发生错误
	//切片存储的元素为字符串中各字符在UTF-8表中的十进制表示，查看字符需要再转换为UTF-8表中十进制对应的字符
	fmt.Printf("转换结果：%v\n对应ASCII码表的字符：%c\n数据类型：%T",slice,slice,slice)
}


/**
切片转字符串
str := string()函数将切片转为字符串，括号里写要转为字符串的切片，切片的数据类型为byte
切片元素为十进制UTF-8表的整数，转换后的字符串为UTF-8表整数对应的字符
 */
func case_4(){
	str := string([]byte{104,101,108,108,111,32,103,111,108,97,110,103,33}) //str接收转换后的字符串
	fmt.Printf("转换结果：%v    数据类型：%T",str,str)
}


/**
十进制转换为其他进制
strconv.FormatInt()函数将十进制转为其他进制  括号里两个参数:第一个参数为要转的十进制数，第二个参数为要转的进制
函数返回值一个,为转换结果
 */
func case_5(){
	var number int64 = 565643
	base_2 := strconv.FormatInt(number,2)
	base_8 := strconv.FormatInt(number,8)
	base_16 := strconv.FormatInt(number,16)
	fmt.Print("十进制：",number,"\n转为二进制：",base_2,"\n转为八进制：",base_8,"\n转为十六进制：",base_16)
}

/**
字符串转为大写或小写
strings.ToLower()函数：字符串转为小写  一个参数：要转的原字符串  一个返回值：转换后的新字符串
strings.ToUpper()函数：字符串转为大写  一个参数：要转的原字符串  一个返回值：转换后的新字符串
*/
func case_6(){
	str := "hello golang!"
	fmt.Println(strings.ToLower(str)) //转为小写
	fmt.Println(strings.ToUpper(str)) //转为大写
}