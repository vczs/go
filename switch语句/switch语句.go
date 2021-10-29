package main

import "fmt"

func main()  {
	case_1()
	case_2()
}

/**
go中switch语句不需要使用break，某个case后语句执行完毕会自动跳出switch
 */
func case_1(){
	var number int8 = 1
	switch number {
	case 0 :fmt.Println("0")
	case 1 :fmt.Println("1")
	case 2 :fmt.Println("2")
	default:fmt.Println("default")
	}
}


/**
go中switch可以不屑条件变量，在case后加条件进行判断，和if else语句相似
 */
func case_2(){
	var number int8 = 2
	switch {
	case number == 0 :
		fmt.Println("number == 0")
	case number == 1 || number < 2:
		fmt.Println("number == 1 || number < 2")
	case number == 2 && number > 0:
		fmt.Println("number == 2 && number > 0")
	default:
		fmt.Println("default")
	}
}