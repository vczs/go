package main

// 注意：运行该项目之前先设置 GO111MODULE=off 环境，并且将该项目放到GOPATH/src目录下运行！！！

/**
包的引进：go的类中可以引进GOPATH/src目录下的包，在import后输入要引进的包的相对路径即可
注意：只能引进本机环境变量中的go工作项目地址的src目录下的包！！！
 */
import (
	"fmt"
	"go-base/导包/tool"
)

func init(){//go的类如果有init()函数会优先执行init()函数
	fmt.Println("主程序init()函数")
}

func main(){
	fmt.Println(tool.Name)             //访问引进的"tool"中的变量Name,因为引进的"tool"前面自定义了包名，所以只能用自定义的包名访问其成员
	fmt.Println(tool.NumberSum(10,20)) //访问引进的"tool"中的函数NumberSum
}