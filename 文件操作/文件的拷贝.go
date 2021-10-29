package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	case_9()
	case_10()
}

/**
文本文件的拷贝
 */
func case_9(){

	readFilePath := "A:/hello.txt" //要复制的文件路径
	writeFilePath := "C:/hello.txt" //目标文件路径

	//先读取要拷贝的文件
	//file的Open和Close已经被封装到ioutil.ReadFile函数内，所以我们不用对file进行打开和关闭操作，直接读取即可。
	data , readErr:= ioutil.ReadFile(readFilePath)
	if readErr != nil {
		fmt.Printf("文件读取错误：%v",readErr)
		return
	}

	//将读取到的内容写入到目标文件即完成了文件的拷贝
	//file的Open和Close已经被封装到ioutil.WriteFile()函数内，所以我们不用对file进行打开和关闭操作，直接读取即可。
	writeErr := ioutil.WriteFile(writeFilePath,data,0)
	if writeErr != nil {
		fmt.Printf("文件写入失败：%v",writeErr)
		return
	}

	fmt.Println("文件拷贝成功！！！")
}


/**
非文本文件的拷贝（视频、图片、音乐）
此方法适用于所有视频、图片、音乐文件的拷贝，因为他们都是二进制文件
 */
func case_10(){
	srcFilePath := "A:/mv.mp4"
	dstFilePath := "C:/视频.mp4"
	written , err := CopyFile(srcFilePath,dstFilePath)
	if err != nil || written == 0 {
		fmt.Printf("文件拷贝失败 err=%v",err)
	}else{
		fmt.Printf("总拷贝大小：%v 字节\n文件拷贝成功！！！",written)
		fmt.Print("\n***********拷贝文件结束***********")
	}
}
func CopyFile(srcFilePath string , dstFilePath string) (written int64 , err error) {
	fmt.Println("***********拷贝文件开始***********")

	//打开源文件获取文件读取流指针
	srcFileInfo , openErr := os.Open(srcFilePath)
	if openErr != nil {
		fmt.Printf("源文件打开失败 openErr=%v",openErr)
		return
	}
	fmt.Println("源文件已打开。。。")
	defer srcFileInfo.Close() //函数执行完要及时关闭源文件
	defer fmt.Println("源文件已关闭。。。")
	reader := bufio.NewReader(srcFileInfo) //文件读取流指针

	//打开目标文件获取文件写入流指针
	datFileInfo , openFileErr := os.OpenFile(dstFilePath , os.O_WRONLY | os.O_CREATE,0 )
	if openFileErr != nil {
		fmt.Printf("目标文件打开失败 openFileErr=%v",openErr)
	}
	fmt.Println("目标文件已打开。。。")
	defer srcFileInfo.Close() //函数执行完要及时关闭目标文件
	defer fmt.Println("目标文件已关闭。。。")
	writer := bufio.NewWriter(datFileInfo)//文件写入流指针

	//用GO自带包io的Copy()方法进行文件的拷贝，将获取到的 读取流指针 和 写入流指针 传给他即可
	return io.Copy(writer,reader)
}