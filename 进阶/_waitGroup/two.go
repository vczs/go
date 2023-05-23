package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	fmt.Println("当前value=", c.count)
}

func Worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done() // 执行后WaitGroup的计数值-1
	time.Sleep(time.Second)
	c.Increase()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	wg.Add(10) // 设置WaitGroup的起始值

	for i := 0; i < 10; i++ {
		go Worker(&counter, &wg)
	}

	// wg.Wait()：goroutine直阻塞 直到WaitGroup的计数值变为0 才执行后面的代码
	wg.Wait()

	fmt.Println("所有工作完成,count最终结果:", counter.count)
}
