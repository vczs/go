package main

import "fmt"

func main() {
	case_13()
}

/**
遍历字符串
将字符串转为切片存入切片中，遍历切片
*/
func case_13(){
	str := "你好 golang!"
	slice := []rune(str) //字符串转为切片
	fmt.Println("slice：",slice,len(slice),cap(slice))
	for i := 0 ; i < len(slice) ; i++ {
		fmt.Printf("%c\n",slice[i]) //切片存储的元素为字符串中各字符在UTF-8表中的十进制表示，遍历原字符要将各十进制元素再转为在UTF-8表中字符表示形式
	}
}
