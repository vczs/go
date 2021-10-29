package main

import "fmt"

func main() {
	case_4()
	case_5()
}

/**
管道的关闭
使用close(管道名)方法对管道进行关闭，管道关闭后只能取出数据不能放入数据
 */
func case_4(){
	intChan := make(chan int,3)
	intChan <- 20
	intChan <- 30
	close(intChan)//关闭intChan管道
	numTem := <- intChan //管道关闭后还是可以取出数据的
	fmt.Printf("numTem=%v",numTem)
}

/**
管道的遍历
 */
func case_5(){
	intChan := make(chan int,50)
	//给管道放入50个int数据
	for i := 1 ; i <= 50 ; i++ {
		intChan <- i
	}
	//管道遍历前要先关闭管道 不然当管道内数据取完后会报错
	close(intChan) //不关闭管道for range将管道内数据取完后还会接着取，因为数据已取完，再取就会报错。
	//使用for range对管道进行遍历
	for v := range intChan {
		fmt.Printf("v=%v  ",v)
	}
}