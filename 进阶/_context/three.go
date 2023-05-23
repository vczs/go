package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ctx := context.WithValue(context.Background(), "one", "go语言")
	go g1(ctx, &wg)
	wg.Wait()
}

func g1(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("g1:", ctx.Value("one"))
	go g2(context.WithValue(ctx, "two", "hello word！"), wg)
	wg.Done()
}

func g2(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("g2:", ctx.Value("one"), ctx.Value("two"))
	wg.Done()
}
