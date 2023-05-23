package main

import (
	"fmt"
	"time"
)

func main(){
	case_6()
	case_7()
}

/**
开启一个writeData协程，向intChan管道写入50个int类型数据
开启一个readData协程，从intChan管道里读取writeData写入的50个int类型数据
主线程等待writeData协程和readData协程都执行完毕再退出
*/
//向intChan写入50个数据
func writeData(intChan chan int){
	for i := 1 ; i <= 50 ; i++ {
		intChan <- i
		fmt.Printf("放入数据：%v\n",i)
	}
	close(intChan)//数据写完就关闭intChan，防止后面数据取完再取会报错，关闭intChan不影响intChan里的数据取出
}
//取出intChan里的50个数据，取完向exitChan放入一个数据
func readData(intChan chan int , exitChan chan bool){
	for v := range intChan{
		fmt.Printf("取出数据：%v\n",v)
	}
	exitChan <- true //intChan里的数据取完，向exitChan放入一个数据，通知主线程退出
	close(exitChan)//数据写完就关闭exitChan，防止后面数据取完再取会报错，关闭exitChan不影响exitChan里的数据取出
}
func case_6(){
	intChan := make(chan int,50) //intChan存取50个int类型数据
	exitChan := make(chan bool,1) //exitChan存储是否退出程序的bool类型数据
	go writeData(intChan) //开启writeData协程放入数据
	go readData(intChan,exitChan) //开启writeData协程取出数据
	//遍历exitChan
	for v := range exitChan {
		//如果exitChan里没有数据，程序会一直阻塞直至读取到数据
		if v {
			fmt.Printf("主线程退出：%v",v)
			break//读取到数据并且为true时退出for循环
		}
	}
	//for循环退出 主线程结束 程序结束退出
}


/**
使用管道加协程，统计1-1000的整数中，有哪些素数。
 */
//创建1-800的整数
func writeNum(intChan chan int){
	for i := 2 ; i < 1000 ; i++{ //1不是素数，所以从整数2开始存入intChan
		intChan <- i //将整数i放入intChan管道
	}
	close(intChan)//数据写完就关闭intChan，防止后面数据取完再取会报错，关闭intChan不影响intChan里的数据取出
}
//从intChan管道中统计素数，如果是素数就将其放入resChan管道
func findPrime(intChan chan int,resChan chan int,exitChan chan bool){
	var flag bool //num是否为素数的标识
	for {
		flag = true //先假设num是素数
		num , ok := <- intChan //从intChan取出整数num
		if !ok { //管道中数据取完或管道中取不出数据
			exitChan <- true //通知主线程执行完毕
			break //跳出for循环，协程结束退出
		}
		for i := 2 ; i < num ; i++ { //判断num是不是素数
			if num % i == 0 {
				flag = false //num不是素数
				break //不是素数直接跳出“判断num是不是素数”的for循环，去判断下一个
			}
		}
		if flag {
			resChan <- num//flag为真,说明num是素数,就将该num放入resChan管道中
		}
	}
}
func case_7(){
	start := time.Now().UnixNano()
	intChan := make(chan int,500) //存放1-800的整数
	resChan := make(chan int,500) //存放1-800里的统计的素数结果
	exitChan := make(chan bool,8) //存放协程是否执行完毕

	go writeNum(intChan) //开启创建整数协程

	for i := 0 ; i < 8 ; i++ {
		go findPrime(intChan,resChan,exitChan) //开启八个从intChan管道中统计素数的协程
	}

	go func(){ //开启一个匿名函数协程，统计“统计素数”的协程执行情况
		count := 0
		for {
			if count == 8 {
				break //如果开启的所有协程都执行完毕就跳出for range遍历
			}
			<- exitChan //取出的数不需要，直接丢弃
			count++ //增加“执行完毕”的协程数量
		}
		close(resChan) //统计结束，数据写完就关闭resChan，防止后面数据取完再取会报错，关闭resChan不影响resChan里的数据取出
		close(exitChan) //统计结束，关闭exitChan协程执行情况管道
	}()

	//遍历resChan统计素数结果管道,并统计一共有多少个素数
	count := 0
	fmt.Print("素数有：{")
	for v := range resChan {
		count++
		fmt.Printf("%v ",v)
	}
	fmt.Printf("}共%v个。",count)

	end := time.Now().UnixNano()
	fmt.Printf("\n耗时=%v纳秒", end - start)
}