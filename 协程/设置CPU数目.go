package main

import (
	"fmt"
	"runtime"
)

func main(){
	cpuNum := runtime.NumCPU() //获取当前设备CPU数目
	fmt.Printf("cpuNum = %v\n数据类型：%T\n",cpuNum,cpuNum)

	runtime.GOMAXPROCS(4)//设置当前程序使用多少个CPU
	fmt.Println("设置成功。")
}
