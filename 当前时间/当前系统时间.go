package main

import (
	"fmt"
	"time"
)

func main() {
	case_1()
	case_2()
	case_3()
}

func case_1() {
	times := time.Now()
	fmt.Printf("当前时间：%v年%v月%v日%v时%v分%v秒\n", times.Year(), times.Month(), times.Day(), times.Hour(), times.Minute(), times.Second())
	//go中times.Month()获取的月份时英文月份表示，数据类型为time.Month
	//要将其表示为数字月份需要强制转换为int数据类型
	fmt.Printf("当前时间：%v年%v月%v日%v时%v分%v秒\n", times.Year(), int(times.Month()), times.Day(), times.Hour(), times.Minute(), times.Second())
}

func case_2() {
	times := time.Now() //获取系统当前时间
	//自定义输出格式 注意：时间必须是2006年1月2日15时04分05秒,任何一个不正确都影响输出结果
	fmt.Printf("格式化后时间：%v\n", times.Format("2006/01/02 15:04:05"))
}

func case_3() {
	times := time.Now()
	fmt.Printf("时间戳(秒)：%v\n", times.Unix())      //获取时间戳 单位秒
	fmt.Printf("时间戳(毫秒)：%v\n", times.UnixNano()) //获取时间戳 单位毫秒
}
