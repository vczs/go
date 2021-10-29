package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
}

func case_1() {
	var age int8
	fmt.Println("请输入您的年龄:")
	fmt.Scanln(&age) //获取控制台输入,&赋值符号,将控制台输入赋值给变量age
	if 0 < age && age < 18{
		fmt.Print("未成年")
	}else if age >= 18{
		fmt.Print("成年了")
	}else{
		fmt.Print("输入错误")
	}
}


func case_2() {
	if age := 19 ; age < 18{ //可以在if语句后声明变量并赋值再进行判断,声明变量和判断条件之间加";"即可
		fmt.Print("未成年")
	}else{
		fmt.Print("成年了")
	}
}


func case_3() {
	var year int
	fmt.Println("请输入年份:")
	fmt.Scanln(&year)
	if ( year % 4 == 0 && year % 100 != 0 ) || year % 400 == 0{
		fmt.Print(year,"是闰年")
	}else{
		fmt.Print(year,"不是闰年")
	}
}