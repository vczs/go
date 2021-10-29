package main

import "fmt"

func main(){
	case_1()
	case_2()
}

func case_1(){
	var name string
	var age int8
	fmt.Println("请输入姓名 年龄 工资 每项输入完回车")
	fmt.Scanln(&name)//获取控制台的输入赋值给变量，用&后面跟要赋值给的变量
	fmt.Scanln(&age)
	fmt.Printf("姓名：%v\n年龄：%v",name,age)
}


func case_2(){
	var name string
	var age int8
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 年龄 工资 性别 用空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass) //格式化获取控制台输入
	fmt.Printf("姓名：%v\n年龄：%v\n工资：%v\n性别：%v\n",name,age,sal,isPass)
}