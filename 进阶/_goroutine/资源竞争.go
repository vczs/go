package main

import (
	"fmt"
	"sync"
	"time"
)

/**
使用goroutine完成
计算 1-10 各个数的阶乘，并把各个数的阶乘放入到map中。

通过 lock sync.Mutex 这个全局变量“锁”解决协程资源竞争问题
sync.Mutex是GO自带包sync里的结构体Mutex，它有两个方法：加锁Lock() 和 解锁Unlock()
 */
var (
	myMap = make(map[int]int,10)
	lock sync.Mutex  //lock是一个全局变量的锁，数据类型是：sync.Mutex。  作用：全局互斥锁
)
//计算传入参数n的阶乘，并存入myMap中
func calculation(n int){
	res := 1
	for i := 1 ; i <= n ; i++ {
		res *= n
	}
	lock.Lock() //加锁，防止多个协程使用myMap时出现协程资源竞争问题
	myMap[n] = res
	lock.Unlock()//解锁，使用完myMap就解锁
}
//开启多个协程求 1-10 的阶乘
func main(){
	for i := 1 ; i <= 10 ; i++ {
		go calculation(i)
	}
	time.Sleep(time.Second * 3) //休息3秒，不然主线程过早退出，没执行完的协程也会跟着退出，myMap里的数据会不完整。
	lock.Lock() //加锁，主线程不知道myMap是否使用完毕，有可能会出现资源竞争问题，所以遍历myMap也要加锁
	for index , value := range myMap {
		fmt.Printf("myMap[%v] = %v\n",index,value)
	}
	lock.Unlock()//解锁，使用完myMap就解锁
}