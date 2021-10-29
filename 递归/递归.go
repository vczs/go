package main

import "fmt"

func main(){
	case_1(10)
	case_2(20)
	case_3(30)
	case_4(40)
	case_5(50)
}

/**
递归：输入一个数如果大于2就递归并打印
*/
func case_1(number int) {
	if number > 2 {
		number--
		case_1(number)
	}
	fmt.Println(number)
}

/**
递归：输入一个数如果大于2就递归,小于等于2就打印
*/
func case_2(number int) {
	if number > 2 {
		number--
		case_2(number)
	} else {
		fmt.Println(number)
	}
}

/**
方程：已知f(1)=3;f(n)=2*f(n-1)+1,求f(n)的值
*/
func case_3(n int) int {
	if n == 1 {
		return 3 //当条件中n==1时，返回条件中f(1)=3给出的结果3
	} else {
		return 2*case_3(n-1) + 1 //n未知就返回条件开始递归
	}
}

/**
斐波那契数列
输入斐波那契数列中数的位置，就返回斐波那契数列中该位置对应的数，数列中第一位和第二位的值已知
斐波那契数列：1 1 2 3 5 8....构成了一个序列,这个数列的特点：前面相邻两项之和，构成了后一项。
*/
func case_4(position int) int {
	if position == 1 || position == 2 { //返回已知值
		return 1
	} else {
		return case_4(position-2) + case_4(position-1) //未知就返回条件开始递归
	}
}

/*
有一堆桃子，猴子第一天吃了其中一半，并再多了一个！
以后每天猴子吃其中一半再多吃一个，当吃到第十天（还没吃），发现只有1个桃子了
问：第x(1 <= x <= 10)天还有多少个桃子
*/
func case_5(number int) int {
	if number > 10 || number < 1 {
		return 0 //返回0表示输入错误，不能超过10天也不能小于1天
	}
	if number == 10 {
		return 1 //返回已知值
	} else {
		return (case_5(number+1) + 1) * 2 //未知就返回条件开始递归
	}
}
