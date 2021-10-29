package main

import (
	"errors"
	"fmt"
)

func main(){
	case_1()
	case_2()
}

/**
函数自定义异常信息返回值
 */
func case_1(){
	errResult := case_1_1("hello golang!") //errResult接收函数执行错误信息，如果为nil则未发生错误
	if errResult != nil { //错误信息返回值不为空
		panic(errResult) //panic()：函数抛出一个异常,括号里写自定义异常内容，string数据类型
	}
}
//该函数只有执行发生异常时才会有返回值，返回值为异常原因  没有异常返回值为nil
func case_1_1(name string) error {
	if name == "hello golang"{
		return nil //函数正常执行未发生异常返回nil
	}else{
		return errors.New("发生错误！！！") //errors.New()：自定义异常信息返回值，括号里写自定义内容，string数据类型
	}
}


/**
Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
 */
func case_2(){
	defer func(){
		if err := recover() ; err != nil { //在defer中通过recover捕获这个异常，然后正常处理。
			fmt.Println("发生错误！！！\n错误原因：", err)
		}
	}()
	panic("抛出panic异常") //抛出一个panic的异常
}