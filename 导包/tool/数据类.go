package tool

import "fmt"

var Name string //数据类的全局变量

func init(){ //go的类如果有init()函数会优先执行init()函数
	fmt.Println("数据类init()函数")
	Name = "数据类变量Name"
}