package main

import "fmt"

/**
goto关键字：Go语言的goto语句可以无条件地转移到程序中指定的行
注意：一定要跳转到goto语句后面的代码行，否则会造成死循环
 */
func main(){
	var number int = 3
	fmt.Println("01")
	fmt.Println("02")
	if number > 1 {
		goto tar
	}
	fmt.Println("03")
	fmt.Println("04")
	tar:fmt.Println("05")
}
