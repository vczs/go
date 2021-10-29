package main

import (
	"fmt"
	"net"
)

/**
服务器端
 */
func main() {
	fmt.Println("服务器开始监听。。。")

	listen , listenErr := net.Listen("tcp","0.0.0.0:8888")
	if listenErr != nil {
		fmt.Printf("listenErr=%v\n",listenErr)
		return
	}

	//函数退出后关闭listen
	defer func(){
		CloseErr := listen.Close() //关闭listen
		if CloseErr != nil {
			fmt.Printf("listen CloseErr=%v\n",CloseErr)
			return
		}
		fmt.Println("listen已关闭。。。")
	}()

	for{
		serverConn , serverConnErr := listen.Accept()
		if serverConnErr != nil {
			fmt.Printf("connErr=%v\n",serverConnErr)
		}else{
			ip := serverConn.RemoteAddr().String()
			fmt.Printf("服务器有新的连接：%v   IP地址：%v\n",serverConn,ip)
			go process(serverConn,ip)
		}
	}

}

func process(serverConn net.Conn , ip string){
	//函数退出后关闭serverConn
	defer func(){
		CloseErr := serverConn.Close() //关闭serverConn
		if CloseErr != nil {
			fmt.Printf("serverConn CloseErr=%v\n",CloseErr)
			return
		}
		fmt.Printf("{%v}serverConn已关闭。。。",ip)
	}()
	//循环读取消息
	for{
		buf := make([]byte,1024)//创建一个切片接收读取的消息
		serverReadN , serverReadNErr := serverConn.Read(buf) //返回接收的数据大小 和 错误信息
		if serverReadNErr != nil {
			fmt.Printf("%v客户端断开连接  serverReadErr=%v\n",ip,serverReadNErr)
			return
		}
		//buf存储接收到的内容，buf是byte切片类型所以要强转为string类型
		//接收的数据大小是serverReadN，所以打印buf[0-serverReadN]的内容
		//buf默认大小是1024，如果不设置打印buf[0-serverReadN]的内容，就会将buf的[serverReadN-1024]多余内容打印出来
		fmt.Printf("客户端{%v}：%v\n",ip,string(buf[:serverReadN]))
	}
}