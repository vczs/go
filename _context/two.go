package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	worker(ctx)
}

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("超时了。。。")
	case <-job():
		fmt.Println("任务完成。。。")
	}
}

func job() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}()
	return ch
}
