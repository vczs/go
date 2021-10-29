package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	case_2()
	case_3()
}

/**
带缓冲区的文件读取
缓冲区的作用是：在读取文件时不是一次性将全部文件读取到内存，而是读取一部分处理一部分，这样就可以处理一些大的文件。
缓冲区默认大小为4096
 */
func case_2(){
	//file有三种叫法：1.file对象  2.file指针  3.file文件句柄
	file , openErr := os.Open("A:/test.txt")
	if openErr != nil {
		fmt.Println("打开文件错误：",openErr)
		return
	}
	fmt.Println("*****文件已打开！*****")

	//defer语句后面的代码会在函数执行完毕时最后执行
	defer file.Close() //函数执行完毕时要及时的关闭文件句柄，否则会有内存泄漏
	defer fmt.Print("*****文件已关闭！*****")

	//开始读取文件
	//reader的默认缓冲区大小为4096
	reader := bufio.NewReader(file)//用GO自带包bufio里面的NewReader()方法读取文件，返回一个读取流指针
	//此时reader读取流已经和file文件绑定，通过reader对file进行读取操作
	for {
		//reader读取流指针里有一个ReadString()方法可以读取file文件的内容
		//str为读取到的内容，读取发生错误的信息会返回给err，读取成功时err为nil
		str , err := reader.ReadString('\n') //读到一个换行就结束
		fmt.Print(str)
		if err != nil {
			fmt.Print("\n")
			if err == io.EOF {
				fmt.Printf("文件读取到末尾：%v\n",err)
			}else{
				fmt.Printf("文件读取发生错误：%v\n",err)
			}
			break
		}
	}
}


/**
一次性读取文件：一次将整个文件读取到内存中（不带缓冲区)
使用GO自带包ioutil里的ReadFile()函数进行一次性读取文件
此方法适用于文件不大的情况,文件太大效率会很低！！！
 */
func case_3(){
	//content接收读取到的文件内容，读取发生错误的信息会返回给err，读取成功时err为nil
	file := "A:/test.txt"
	//file的Open和Close已经被封装到ioutil.ReadFile函数内，所以我们不用对file进行打开和关闭操作，直接读取即可。
	content , err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("文件读取发生错误：%v\n",err)
		return
	}
	//读取的内容是一个byte切片，切片元素为file内容的十进制表示（UTF-8表）
	//查看file原内容要将十进制表示转换为对应字符表示（UTF-8表）
	fmt.Println(content) //十进制切片表示
	fmt.Printf(string(content)) //字符表示
}