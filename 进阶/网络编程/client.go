package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
*
客户端
*/
func main() {
	//连接服务端
	clientConn, clientConnErr := net.Dial("tcp", "192.168.31.95:8888")
	if clientConnErr != nil {
		fmt.Printf("clientConnErr=%v", clientConnErr)
		return
	}
	fmt.Printf("连接成功：%v\n", clientConn)

	//函数退出后关闭clientConn
	defer func() {
		clientCloseErr := clientConn.Close() //关闭clientConn
		if clientCloseErr != nil {
			fmt.Printf("clientCloseErr=%v\n", clientCloseErr)
			return
		}
		fmt.Printf("clientConn已关闭。。。")
	}()

	//客户端发送消息
	//获取终端输入的内容的指针
	reader := bufio.NewReader(os.Stdin) //os.Stdin读取终端输入的内容  os.Stdin本质上是一个文件，所以会使用bufio.NewReader()

	//循环向服务器发送消息
	for {
		//从指针里读取 客户端发送消息的文件 内容
		con, conErr := reader.ReadString('\n') //读取到'\n'字符结束
		if conErr != nil {
			fmt.Printf("conErr=%v", conErr)
			return
		}

		conEnd := strings.Trim(con, " \n\r") //去掉con字符串里空格换行并返回给conEnd

		//如果客户端输入“exit”就退出程序
		if conEnd == "exit" {
			fmt.Println("程序已退出。。。")
			return
		}

		//将读取到的内容发送给服务端
		clientWriteN, clientWriteNErr := clientConn.Write([]byte(conEnd)) //将string类型的con转为socket信息管道规定的[]byte类型
		if clientWriteNErr != nil {
			fmt.Printf("nErr=%v", clientWriteNErr)
			return
		}
		fmt.Printf("已发送[%v字节]\n", clientWriteN)
	}
}
