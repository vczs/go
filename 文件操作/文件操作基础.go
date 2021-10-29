package main

import (
	"fmt"
	"os"
)

func main(){
	case_1()
}

/**
文件在程序中以流的形式进行操作
流：文件在数据源（文件）和程序（内存）之间经历的路径
输入流（读取）（input）：数据从数据源到程序的路径
输出流（写入）（output）：数据从程序到数据源的路径
 */

/**
用GO自带的os包下的Open方法打开一个文件
file接收文件指针（文件对象）
err接收打开文件失败信息，打开成功err为nil
 */
func case_1(){
	//file有三种叫法：1.file对象  2.file指针  3.file文件句柄
	file , openErr := os.Open("A:/test.txt")
	if openErr != nil {
		fmt.Println("打开文件错误：",openErr)
		return
	}
	//打开成功就输出 查看文件
	fmt.Println("test:",file) //因为file是一个指针类型，所以输出的是一个地址（文件地址）

	//关闭文件Close()
	CloseErr := file.Close()
	if CloseErr != nil {
		fmt.Println("关闭文件错误：",CloseErr)
	}
}