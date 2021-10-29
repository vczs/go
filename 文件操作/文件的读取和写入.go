package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	case_7()
}

/**
先读取已存在文件的内容并打印在控制台，再给文件追加内容。
 */
func case_7(){
	//打开已存在文件
	filePath := "A:/hello.txt"
	//os.O_RDWR文件读取并写入模式 os.O_APPEND在打开的文件原内容基础上追加 这两个属性均来自于GO自带包os
	//第二个参数为文件打开模式，这个参数传入什么很重要，直接关系到文件写入结果
	file , openErr := os.OpenFile(filePath,os.O_RDWR | os.O_APPEND,0)
	if openErr != nil {
		fmt.Printf("OpenFile err=%v",openErr)
		return
	}
	fmt.Println("*****文件已打开！*****")
	defer file.Close() //函数执行完毕时要及时的关闭文件句柄，否则会有内存泄漏
	defer fmt.Println("*****文件已关闭！*****")

	//读取文件内容并打印
	reader := bufio.NewReader(file)//获取读取流指针
	for {
		str , readerErr := reader.ReadString('\n')//读到一个换行就结束
		fmt.Print(str)
		if readerErr != nil {
			if readerErr == io.EOF {
				fmt.Printf("文件读取到末尾：%v\n",readerErr)
			}else{
				fmt.Printf("文件读取发生错误：%v\n",readerErr)
			}
			break
		}
	}


	//追加写入文件
	writer := bufio.NewWriter(file) //用GO自带包bufio里的NewWriter()方法对文件进行写入操作，返回一个file文件写入流指针
	//循环执行3次该操作
	for i := 0 ; i < 3 ; i++ {
		writer.WriteString("ABCD\n") //file文件写入流指针里有一个WriteString()方法对file文件进行写入
	}
	//因为writer是带缓存的，在调用WriteString()方法时是先写入到缓存中
	//写入操作执行完要执行Flush()方法才会将writer缓存的内容真正写入到file文件中
	//注意：写入操作执行完不执行Flush()方法文件不会file更新
	writer.Flush()
	fmt.Println("文件追加写入成功！！！")
}
