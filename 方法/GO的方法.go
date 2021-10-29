package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
}

/**
go中的方法是和结构体绑定的，绑定了结构体的函数就叫方法（不仅仅是结构体，也可以是其他自定义数据类型）,方法名首字母大写可以在其他包内访问
注意：绑定了结构体的方法，只能通过该结构体数据类型声明的变量来调用，不能直接调用或用其他数据类型声明的变量调用
 */
type Method struct{
	Name string
	Number int
}
func (method Method) case_1_1(number int) bool { // case_1_1()是绑定了Method数据类型的方法，就是说case_1_1()该方法作用于Method类型 method参数结构体名可以自定义，数据类型必须是已创建的结构体
	method.Name = "method"
	method.Number = 104
	number = 2
	fmt.Printf("method的地址%p   method.Name的地址%p\n",&method,&method.Name)
	fmt.Println("case_1_1():",method.Name,method.Number,number)
	return true
}
func case_1(){
	var method Method //声明一个Method数据类型的变量（实例），这里定义的method和case_1_1()中的method不是同一个,调用时会将method值拷贝传给case_1_1()，虽是互相绑定但只是规定谁可以调用
	method.Name = "method~" //给Method生命的变量method的字段赋值
	method.Number = 105
	number := 1
	result := method.case_1_1(number) //调用与其数据类型绑定的方法，调用时会将method之前的赋值（值类型数据）的变量值拷贝过去给case_1_1()方法
	fmt.Printf("method的地址%p   method.Name的地址%p\n",&method,&method.Name) //这里method的地址和case_1_1()中的method地址不同，所以调用case_1_1()方法时也会将method进行值拷贝传给case_1_1()方法
	fmt.Println("case_1():",method.Name,method.Number,number,result) //go中函数调用使用的是值拷贝（值类型数据），所以case_1_1()的赋值不会影响case_1()
}


/**
结构体是值类型，在方法调用中遵守值拷贝传递方式；如果希望在方法中修改结构体变量的值，可以通过绑定结构体指针的方式
 */
type Circle struct {
	radius float32
}
//为了提高效率，我们通常将方法和结构体的指针绑定
func (circle *Circle) case_2_1() float32 {
	(*circle).radius = 6.0 //指针绑定  这里修改的是case_2()中声明的变量的本身
	fmt.Printf("circle的值%p   circle.radius的地址%p\n",circle,&circle.radius) //结构体地址 字段地址和case_2()里的都相同
	//因为circle是指针，所以要使用(*circle)方式访问
	return 3.14 * (*circle).radius * (*circle).radius // (*circle)等价于circle，go的底层做了优化，变量是指针时，方法内会自动访问他指向的值，也就是说自动为其变量加上*
}
func case_2(){
	var circle Circle
	circle.radius = 5.0
	result := circle.case_2_1() //因为是指针绑定，上面的变量传递时不是值拷贝而是变量地址
	fmt.Printf("circle的地址%p   circle.radius的地址%p\n",&circle,&circle.radius) //这里的地址和case_2_1()相同
	fmt.Println("case_2():",circle.radius,result) //在case_2_1()修改的变量的值这里也被修改了
}


/**
方法不仅仅是绑定结构体，也可以是其他自定义数据类型
 */
type integer int
func (number *integer) case_3_1(){
	(*number) += 1
}
func case_3(){
	var number integer
	number = 10
	number.case_3_1()
	fmt.Println("number:",number)
}


/**
如果一个类型实现了String()方法，那么fmt.Println默认会调用这个变量的String()进行输出
 */
type Log struct {
	name string
	flag int
}
func (log *Log) String() string {
	//fmt.Sprintf():用传入的格式化规则符将传入的变量格式化(终端中不会有显示),有一个返回值：返回格式化后的字符串
	str := fmt.Sprintf("自定义方法String()：\nlog.name：%v\nlog.flag：%v",log.name,log.flag)
	return str
}
func case_4(){
	log := Log {
		name : "hello golang!" ,
		flag : 5 ,
	}
	fmt.Println(&log) //因为Log数据类型的log实现了String()方法，所以fmt.Println是调用自定义的方法  自定义方法是指针数据类型，所以要传入地址
}