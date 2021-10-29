package main

import "fmt"

func main(){
	case_8()
}

/**
传统方法在遍历管道时，如果管道没有关闭就会报错。
可以用select方法可以在管道未关闭的情况下遍历管道。
 */
func case_8(){
	intChan := make(chan int,5)
	intChan <- 20
	intChan <- 30
	intChan <- 40
	A:for { //使用for循环一直执行select
		select{
			//如果intChan没有关闭，当这里读取不到数据就会进入下一个case不会报错
			case v := <-intChan:
				fmt.Printf("读取到：%v\n",v)
			default:
				fmt.Printf("读取不到数据，循环结束！！！")
				break A
		}
	}
}