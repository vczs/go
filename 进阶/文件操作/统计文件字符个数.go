package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	case_11()
}

//字符个数统计结构体
type count struct {
	charCount int //字母个数
	numCount int //数字个数
	spaceCount int //空格个数
	otherCount int //其他个数
}

func case_11(){
	var count count //创建结构体
	filePath := "A:/test.txt"
	file , openErr := os.Open(filePath)
	if openErr != nil {
		fmt.Printf("文件打开失败 openErr=%v",openErr)
	}
	defer file.Close() //及时关闭文件

	reader := bufio.NewReader(file) //获取读取流指针

	for {
		str , readErr := reader.ReadString('\n')
		if readErr != nil {
			if readErr == io.EOF {
				fmt.Printf("文件读取到末尾：%v\n",readErr)
			}else{
				fmt.Printf("文件读取发生错误：%v\n",readErr)
			}
			break
		}
		for _ , v := range str {
			fmt.Print(v," ")
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.charCount++
			case v >= '0' && v <= '9' :
				count.numCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			default : count.otherCount++
			}
		}
	}
	//统计结果不兼容中文，如出现中文则会忽略
	fmt.Println("统计结束，统计结果如下：")
	fmt.Printf("字母个数：%v\n数字个数：%v\n空格个数：%v\n其他字符个数：%v",count.charCount,count.numCount,count.spaceCount,count.otherCount)
}
