package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
	case_5()
	case_6()
	case_7()
	case_8()
}

/**
%c:输出单个字符,这里表示该十进制整数在ASCII码中对应的字符
十进制整数97-123对应小写a-z,十进制整数64-90对应大写A-Z
 */
func case_1() {
	for i := 97 ; i < 123 ; i++ {
		fmt.Printf("%c ",i)
	}
	fmt.Println()
	for i := 90 ; i > 64 ; i-- {
		fmt.Printf("%c ",i)
	}
}


func case_2() {
	i := 0
	for {
		if i >= 5 {
			break //跳出最近的for循环语句
		}else{
			fmt.Println(i)
		}
		i++
	}
}


func case_3() {
	str := "hello,golang!"
	for i := 0 ; i < len(str) ; i++ {
		fmt.Printf("%c——%d\n",str[i],str[i]) //%d输出对应的十进制整数
	}
}


func case_4() {
	str := "hello,golang!"
	for index , val := range str { //index为该字符在字符串中的下标,val为该字符
		fmt.Printf("%d——%c——%d\n",index,val,val)
	}
}


func case_5() {
	str := "你好,go语言!"
	for index , val := range str { //go语言中中文占3个字符,for range语句会自动识别中文,index为在字符串中的起始下标
		fmt.Printf("%d——%c——%d\n",index,val,val)
	}
}


func case_6() {
	fmt.Println("100之内9的倍数有：")
	var number int8
	var sum int
	for i := 1 ; i <= 100 ; i++ {
		if( i % 9 == 0) {
			sum += i
			number += 1
			fmt.Printf("%d ",i)
		}
	}
	fmt.Printf("\n共有%d个,总和%d",number,sum)
}


func case_7() {
	var inputScore float32
	var score_sum_1 float32
	var score_sum_2 float32
	var score_sum_3 float32
	var score_sum_temp float32
	var score_sum float32
	for i := 1 ; i < 4 ; i++ {
		for j := 1 ; j < 4 ; j++{
			fmt.Printf("请输入%d班第%d名学生成绩:\n",i,j)
			fmt.Scan(&inputScore)
			score_sum_temp += inputScore
		}
		switch{
		case i == 1 :
			score_sum_1 = score_sum_temp
		case i == 2 :
			score_sum_2 = score_sum_temp
		case i == 3 :
			score_sum_3 = score_sum_temp
		}
		score_sum += score_sum_temp
		score_sum_temp = 0.0
	}
	fmt.Printf("一班的平均成绩：%v\n",score_sum_1/5)
	fmt.Printf("二班的平均成绩：%v\n",score_sum_2/5)
	fmt.Printf("三班的平均成绩：%v\n",score_sum_3/5)
	fmt.Printf("总成绩：%v,班级平均成绩：%v",score_sum,score_sum/3)
}


/**
有一笔钱，当大于等于50000时每次减少1000块，反之每次减少其总钱数的百分之5
问：一共可以这样进行多少次
 */
func case_8() {
	money := 100000.0
	count := 0
	for {
		if money <= 0{
			break
		}
		count++
		if money <= 50000{
			money -= 1000
			continue
		}
		money -= (money*0.05)
	}
	fmt.Println(count)
}