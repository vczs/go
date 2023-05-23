package main

import "fmt"

/**
管道(channel)：的本质就是一个数据结构 - 单向队列。

特点：
1.先进先出（FIFO:first in first out）
2.线程安全，多goroutine访问时，不会发生资源竞争问题，无需加锁，就是说channel本身就是线程安全的。

注意事项：
1.channel是指针类型，必须初始化才能完成数据写入，即必须make后才能使用。
2.channel是有数据类型的，一个string数据类型的channel只能存放string类型数据。
3.make时不设置容量或者设置为false就是无缓冲channel，设置大于0的容量就是有缓冲channel。
4.channel关闭后不能再放数据，再放会报错。

在没有使用协程的情况下：
1.当channel内放满时再放就会报错，长度不能超过容量；当channel内的数据全部取出时(长度等于零)再取就会报错。
2.当channel内放满时再放就会报错，但将管道内数据取出后，仍可以继续放入；取出几个数据便可以再放入几个数据，直至再次放满。
3.使用for range对未关闭的管道进行遍历会报错。

channel声明:
var ch chan 数据类型
管道可以声明为只读或只写（管道在默认情况下是双向的）
var ch chan<- 数据类型 (只写管道)（管道只能放入数据，不能取出）
var ch <-chan 数据类型 (只读管道)（管道只能取出数据，不能放入）
ch = make(chan 数据类型 , 容量)		// 管道必须make后才能使用
*/

func main() {
	case_1()
	case_2()
}

/**
channel声明和数据存取
*/
func case_1() {
	//管道声明和实例化
	var intChan chan int = make(chan int, 3) //intChan必须make后才能使用，chan int是intChan的数据类型，3是intChan的容量
	//向管道内放入数据
	intChan <- 10                                                            //将10放入管道
	intChan <- 12                                                            //将12放入管道
	fmt.Printf("intChan的长度：%v  intChan的容量：%v\n", len(intChan), cap(intChan)) //查看管道的长度和容量
	//从管道内取出数据
	var num1 int
	num1 = <-intChan
	num2 := <-intChan
	fmt.Printf("num1=%v   num2=%v\n", num1, num2)
	fmt.Printf("intChan的长度：%v  intChan的容量：%v\n", len(intChan), cap(intChan)) //查看管道的长度和容量
}

func case_2() {
	//管道声明和实例化
	var intChan chan int = make(chan int, 2)
	//向管道内放入数据
	intChan <- 25                                                            //将25放入管道
	intChan <- 15                                                            //将15放入管道
	fmt.Printf("intChan的长度：%v  intChan的容量：%v\n", len(intChan), cap(intChan)) //查看管道的长度和容量
	//当管道放满时，再放就会报错。
	<-intChan                                                                //将intChan管道内的数据取出来不要了，给管道清理个位置来放入需要使用的数据
	fmt.Printf("intChan的长度：%v  intChan的容量：%v\n", len(intChan), cap(intChan)) //查看管道的长度和容量
	intChan <- 18                                                            //但将管道内数据取出后，仍可以继续放入，取出几个数据便可以再放入几个数据，直至再次放满。
	fmt.Printf("intChan的长度：%v  intChan的容量：%v\n", len(intChan), cap(intChan)) //查看管道的长度和容量
}
