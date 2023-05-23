package main

import (
	"fmt"
	"time"
)

/**
当开启协程出现panic时，如果程序员没有捕获这个panic，整个程序就会崩溃
如果主线程起的一个协程发生panic，那么主线程和主线程起的所有协程都将崩溃
使用recover()捕获协程的panic，可以防止程序崩溃；当主线程起的一个协程发生panic时，可以保证主线程和主线程起的所有协程继续执行。
 */
func main(){
	//看warn()协程发生的异常是否会影响到hello()协程的执行
	go hello()
	go warn()
	time.Sleep(time.Second * 3) //主线程休眠3秒等待协程执行
}
//打印10句hello golang！
func hello(){
	for i := 0 ; i < 10 ; i++ {
		fmt.Println("hello golang！")
	}
}
//使用defer + recover来捕获函数的异常
func warn(){
	//defer后的语句待函数执行完毕再执行
	defer func(){
		//捕获当前函数发生的异常
		if err := recover() ; err != nil {
			fmt.Printf("warn()发生异常：%v\n",err)
		}
	}()
	panic("warn()抛出一个恐慌！！！") //抛出一个异常
}
