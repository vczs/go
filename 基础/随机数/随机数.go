package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	case_1()
	case_2()
}

func case_1(){
	//go中生成随机数时先要给math的rand中Seed一个当前系统时间
	//系统时间秒和毫秒都可以，毫秒随机生成的更快一些
	rand.Seed(time.Now().Unix())
	result := rand.Intn(100)//获取100以内的随机数,不包括0和括号里的数（0到该数之内）
	fmt.Println(result)
}

func case_2(){
	var count int
	for{
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100)-1
		fmt.Println(num)
		count++
		if num == 0 {
			break
		}
	}
	fmt.Printf("共%d次生成随机数5",count)
}