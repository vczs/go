package main

import (
	"fmt"
	"os"
)

func main(){
	case_8()
}

/**
判断文件或目录是否存在
 */
func case_8(){
	file := "A:/test.txt"
	fileInfo , err := os.Stat(file)
	if os.IsNotExist(err) {
		//用GO自带包os里的函数IsNotExist()判断err为true的话说明文件不存在
		//文件不存在的情况下：fileInfo为nil err不为nil os.IsNotExist(err)为true
		fmt.Printf("文件：%v 不存在：%v",fileInfo,err)
		return
	}
	if err == nil {
		//如果err为nil说明文件存在
		//文件存在的情况下：fileInfo不为nil err为nil os.IsNotExist(err)为false
		fmt.Printf("文件：%v 存在：%v",fileInfo,err)
		return
	}
	//以上无法判断就告诉用户无法确定文件是否存在
	//文件不确定是否存在的情况下：fileInfo为nil err不为nil os.IsNotExist(err)为false
	fmt.Printf("文件：%v 不确定是否存在：%v",fileInfo,err)
}