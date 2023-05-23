package main

import(
	"fmt"
	"strings"
)

/**
闭包：闭包就是一个函数和与其相关的引用环境组合的一个整体
 */
func main(){
	case_1()
	case_2()
}

func case_1() {
	//这里f用来接收case_1_1()函数返回值  因为返回值类型是一个匿名函数 所以当前f是一个匿名函数 func (x int) int
	f := case_1_1() //f接收一个闭包  case_1_1()返回一个闭包
	fmt.Println(f(1)) //f(1) 相当于调用匿名函数func (x int) int 传入参数1  调用闭包f传入参数1
	fmt.Println(f(2))
	fmt.Println(f(3))
}
//累加器
// func (x int) int是case_79_1()函数的返回值类型  因为返回值类型func(x int)int没有函数名，所以是一个匿名函数
func case_1_1() func (int) int { //函数返回值是一个匿名函数
/*******************************以下内容形成一个闭包***********************************/
	var n int = 10
	return func (number int) int {
		n = n + number
		return n //匿名函数返回一个int类型数据
	}
	//这里可以这么理解：闭包是类 返回值匿名函数是操作 n是字段  匿名函数和n构成一个闭包
	//n只在调用函数时初始化一次 当我们反复调用闭包时 n是不断累加的 f(1)执行过后n就更新为11  f(2)执行过后n就更新为13(11+2)
	//要弄明白闭包 就要分析出匿名函数引用到了哪些变量 因为匿名函数和它引用到的变量构成一个闭包
	//n只在调用函数时初始化一次 调用闭包时就不会再初始化了 (变量不会再初始化 上次操作执行后n的值最后是多少 下次调用闭包时n初始值就是多少)
/******************************以上内容形成一个闭包**************************************/
}


/**
编写一个函数传入指定文件后缀名 返回一个闭包
调用闭包 传入一个文件名 如果该文件的后缀名是 以函数传入的指定文件后缀名 结尾就返回原文件名
如果不是就给追加 函数传入的指定文件后缀名 再返回追加后的文件名
 */
func case_2() {
	s := makeSuffix(".jpg") //给函数传入指定文件后缀名.jpg  函数返回一个闭包s
	fmt.Println(s("aaa.jpg")) //调用闭包 给它传入aaa.jpg
	fmt.Println(s("bbb"))
	fmt.Println(s("ccc.avi"))
}
//makeSuffix(suffix string)函数返回值类型为 func (string) string
func makeSuffix(suffix string) func (string) string {
	//返回一个匿名函数
	return func (name string) string {
		endName := "nil"
		if strings.HasSuffix(name,suffix) { //strings.HasSuffix(name,suffix):判断name是否以suffix结尾  是返回true 不是返回false
			endName = name
		}else{
			endName = name + suffix //给name追加suffix
		}
		return endName
	}
}
//因为返回的匿名函数func (name string) string引用到了suffix   所以返回的匿名函数和函数makeSuffix(suffix string)的suffix构成一个闭包
//使用传统方法也可以实现此功能 但是传统方法每次都需要传入后缀名.jpg 而闭包因为可以保留上次引用的值 所以我们传入一次就可以重复使用