package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	case_1()
	case_2()
}

/**
获取命令行参数
该方法在命令提示符运行程序时才可用
 */
func case_1(){
	fmt.Println("命令行参数有：",len(os.Args))
	for index , value := range os.Args {
		fmt.Printf("第%v个参数：%v\n",index+1,value)
	}
}

/**
获取命令行指定名称的参数
该方法在命令提示符运行程序时才可用
 */
func case_2(){
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","admin","用户名")
	flag.StringVar(&pwd,"pwd","a1b2c3","密码")
	flag.StringVar(&host,"h","127.0.0.1","主机IP")
	flag.IntVar(&port,"port",8888,"端口号")

	flag.Parse() //解析变量  这一行一定要有  不解析命令行输入的参数无法获取到

	fmt.Printf("用户名=%v\n密码=%v\n主机IP=%v\n端口号=%v\n",user,pwd,host,port)
}