package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type myWaitGroup struct {
	sync.WaitGroup
	t int64
}

func newMyWaitGroup() *myWaitGroup {
	return &myWaitGroup{}
}

func (w *myWaitGroup) Done(start time.Time) {
	defer w.WaitGroup.Done()
	defer func() {
		w.t += int64(time.Since(start).Seconds())
	}()
}

func myJob(wg *myWaitGroup, start time.Time) {
	t := time.Second * time.Duration(rand.Intn(4))
	defer wg.Done(start)
	time.Sleep(t)
	fmt.Println("运行了", t, "秒")
}

func main() {
	start := time.Now()
	wg := newMyWaitGroup()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go myJob(wg, start)
	}
	wg.WaitGroup.Wait()
	println("一共运行了", wg.t, "秒")
}
