package main

import (
	"fmt"
	"sync"
)

// 两个协程，要求实现交替打印的效果，从1开始打印到20
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch1 := make(chan bool, 0)
	ch2 := make(chan bool, 0)

	go func() {
		defer wg.Done()
		for i := 1; i < 20; i += 2 {
			fmt.Println("第1个协程打印:", i)
			ch2 <- true
			<-ch1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i < 21; i += 2 {
			<-ch2
			fmt.Println("第2个协程打印:", i)
			ch1 <- true
		}
	}()

	wg.Wait()
	fmt.Println("打印完成。。。")
}
