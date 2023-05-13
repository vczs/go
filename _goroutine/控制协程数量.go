package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 每次发放五个协程，所有协程执行完，再发放五个。
func main() {
	ch1 := make(chan struct{}, 5)
	ch2 := make(chan struct{}, 5)

	go sendGoroutine(ch1, ch2)

	for {
		go job(ch2)
		<-ch1
	}
}

func job(ch2 chan struct{}) {
	fmt.Println(rand.Intn(100))
	ch2 <- struct{}{}
}

func sendGoroutine(ch1, ch2 chan struct{}) {
	n := 0
	for range ch2 {
		if n%5 == 0 {
			fmt.Println("开始发放。。。")
			for i := 0; i < 5; i++ {
				ch1 <- struct{}{}
			}
			time.Sleep(time.Second * 2)
		}
		n++
	}
}
