package main

import (
	"fmt"
	"time"
)

/**
主线程每隔一秒输出一句“Hello World!”
在主线程中开启一个协程，协程每隔一秒输出一句“你好，世界！”。
 */
func hello_cn(){
	for i := 1 ; i <= 10 ; i++{
		fmt.Println("你好，世界！----",i)
		time.Sleep(time.Second)
	}
}
func hello_en(){
	for i := 1 ; i <= 10 ; i++{
		fmt.Println("Hello World!----",i)
		time.Sleep(time.Second)
	}
}
func main(){
	go hello_cn()
	hello_en()
}