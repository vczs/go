package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	case_5()
	case_6()
}

/**
打开一个存在的文件，覆盖原内容替换为10句：“你好,GO!”
 */
func case_5(){
	//打开已存在文件
	filePath := "A:/hello.txt"
	//file接收文件指针，err接收发生错误时的信息，未发生错误err为nil
	//os.O_WRONLY文件只写模式 os.O_TRUNC在打开文件同时清空文件内容 这两个属性均来自于GO自带包os
	//第二个参数为文件打开模式，这个参数传入什么很重要，直接关系到文件写入结果
	file , err := os.OpenFile(filePath,os.O_WRONLY | os.O_TRUNC,0)
	if err != nil {
		fmt.Printf("OpenFile err=%v",err)
		return
	}
	fmt.Println("*****文件已打开！*****")

	defer file.Close() //函数执行完毕时要及时的关闭文件句柄，否则会有内存泄漏
	defer fmt.Println("*****文件已关闭！*****")

	//写入文件
	writer := bufio.NewWriter(file) //用GO自带包bufio里的NewWriter()方法对文件进行写入操作，返回一个file文件写入流指针
	//循环执行10次该操作
	for i := 0 ; i < 10 ; i++ {
		writer.WriteString("你好，GO!\n") //file文件写入流指针里有一个WriteString()方法对file文件进行写入
	}
	//因为writer是带缓存的，在调用WriteString()方法时是先写入到缓存中
	//写入操作执行完要执行Flush()方法才会将writer缓存的内容真正写入到file文件中
	//注意：写入操作执行完不执行Flush()方法文件不会file更新
	writer.Flush()
	fmt.Println("文件写入成功！！！")
}


/**
打开一个存在的文件，在原内容的基础上追加5局“hello golang！”
 */
func case_6(){
	//打开已存在文件
	filePath := "A:/hello.txt"
	//os.O_WRONLY文件只写模式 os.O_APPEND在打开的文件原内容基础上追加 这两个属性均来自于GO自带包os
	//第二个参数为文件打开模式，这个参数传入什么很重要，直接关系到文件写入结果
	file , err := os.OpenFile(filePath,os.O_WRONLY | os.O_APPEND,0)
	if err != nil {
		fmt.Printf("OpenFile err=%v",err)
		return
	}
	fmt.Println("*****文件已打开！*****")

	defer file.Close() //函数执行完毕时要及时的关闭文件句柄，否则会有内存泄漏
	defer fmt.Println("*****文件已关闭！*****")

	//追加写入文件
	writer := bufio.NewWriter(file) //用GO自带包bufio里的NewWriter()方法对文件进行写入操作，返回一个file文件写入流指针
	//循环执行10次该操作
	for i := 0 ; i < 5 ; i++ {
		writer.WriteString("hello golang！\n") //file文件写入流指针里有一个WriteString()方法对file文件进行写入
	}
	//因为writer是带缓存的，在调用WriteString()方法时是先写入到缓存中
	//写入操作执行完要执行Flush()方法才会将writer缓存的内容真正写入到file文件中
	//注意：写入操作执行完不执行Flush()方法文件不会file更新
	writer.Flush()
	fmt.Println("文件追加写入成功！！！")
}