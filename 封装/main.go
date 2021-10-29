package main

import (
	"go-base/封装/model"
	"fmt"
)

// 注意：运行该项目之前先设置 GO111MODULE=off 环境，并且将该项目放到GOPATH/src目录下运行！！！

func main(){
	bank := model.NewBank("golang","666666")
	fmt.Println(*bank)
	bank.SetBalance(500.00)
	fmt.Println(bank.GetBalance())
	fmt.Println(*bank)
}
