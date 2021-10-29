package main

import "fmt"

/**
继承：go中继承是一个结构体嵌入匿名结构体的方法实现
可以让某个类型的对象获得另一个类型的对象的属性的方法
它可以使用现有类的所有功能，并在无需重新编写原来的类的情况下对这些功能进行扩展。

多重继承：一个结构体嵌套了多个（两个或两个以上）匿名结构体
为了保证代码的简洁性，建议尽量不要使用多重继承
*/

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
	case_5()
	case_6()
	case_7()
}

/**
先定义一个有共同属性(相同字段、方法)的结构体
 */
type Student struct {
	name string
	score int
}
//绑定共同属性的方法是共同方法，即继承该结构体的对象都需要实现的方法
func (student *Student) ShowInfo(){
	fmt.Printf("姓名：%v   成绩：%v",(*student).name,(*student).score)
}
func (student *Student) setScore(score int){
	(*student).score = score
}


//要继承Student的结构体
type pupil struct {
	Student //pupil嵌入匿名结构体Student，此时pupil具有Student的所有字段和方法
}
//特有方法或字段(Student中没有的)要自己实现，并绑定自己
func (pupil *pupil)testing(){
	fmt.Println("pupil testing...")
}
/**
当结构体和匿名字段有相同字段或方法时，编译器会采用就近访问原则
如想访问匿名结构体字段，可以通过匿名结构体名来区分
*/
func case_1(){
	pupil := &pupil{}
	(*pupil).Student.name = "golang"
	(*pupil).testing()
	(*pupil).Student.setScore(60)
	(*pupil).Student.ShowInfo()
}


//要继承Student的结构体
type graduate struct {
	Student //graduate嵌入匿名结构体Student，此时graduate具有Student的所有字段和方法
	end bool //非公有字段要自己添加
}
//特有方法或字段(Student中没有的)要自己实现，并绑定自己
func (graduate *graduate)testing(){
	fmt.Println("graduate testing...")
}
//graduate也绑定了和Student共有结构体绑定的ShowInfo()方法，方法名相同
func (graduate *graduate) ShowInfo(){
	fmt.Printf("姓名：%v   成绩：%v   end：%v",(*graduate).name,(*graduate).score,(*graduate).end)
}
/**
如果结构体和匿名结构体无相同字段或方法，就可以简化访问
编译器先会找声明的对象对应结构体的字段和绑定的方法，没有的话就会去嵌入的匿名结构体里找字段或与匿名结构体绑定的方法，直到找不到该字段或方法
即无需通过匿名结构体名，可以直接访问匿名结构体里的方法和字段
但是如果有和匿名结构体相同的字段时，就需要通过匿名结构体名来区分，因为编译器是就近访问原则，不区分就无法访问匿名结构体里的字段和方法
 */
func case_2(){
	graduate := &graduate{}
	(*graduate).name = "hello golang" //结构体和匿名结构体无相同字段，可以简化访问
	(*graduate).end = true
	(*graduate).testing()
	(*graduate).setScore(80) //结构体和匿名结构体无相同方法，可以简化访问
	(*graduate).ShowInfo() //结构体和匿名结构体有相同方法，采用就近访问原则，先访问graduate里的该访问，没有才会去嵌入的匿名结构体里找
}


/**
结构体嵌入两个或多个结构体
在两个匿名结构体都有相同的字段和方法（同时结构体本身没有该相同的字段或方法）
在访问时，如果不指明哪个结构体，即不通过匿名结构体名访问该字段或方法，直接简化访问，编译器就会报错
 */
type case_3_1 struct {
	name string
	age int
	id string
}
type case_3_2 struct {
	name string
	score int
	id string
}
//case_91_3嵌入case_91_1和case_91_2两个匿名结构体，这里就使用了多重继承
type case_3_3 struct {
	case_3_1
	case_3_2
	id string
}
func case_3(){
	case_3_3 := case_3_3{}
	case_3_3.case_3_1.name = "hello" //必须通过嵌入匿名结构体名来访问，直接简化访问会报错
	case_3_3.case_3_1.name = "golang"
	case_3_3.age = 20             //嵌入的两个或多个结构体不相同的字段，可以直接简化访问
	case_3_3.id = "hello golang!" //嵌入的两个或多个结构体有相同字段，同时自身结构体也有，会采用就近访问原则只访问自身该相同字段，嵌入的结构体不会打架，所以可以直接简化访问，不会报错
}


/**
如果一个结构体嵌入了有名结构体，那么这种模式就称为组合
访问嵌入的结构体字段或方法时，就必须通过该结构体名访问，不能直接简化访问，会报错
 */
type case_4_1 struct {
	name string
}
type case_4_2 struct {
	a case_4_1 //嵌入了有名结构体  a是名 case_92_1是数据类型
}
func case_4() {
	case_4_2 := case_4_2{}
	case_4_2.a.name = "golang" //嵌入有名结构体时访问其字段必须通过
}


/**
嵌入匿名结构体后，可以在常见变量时，直接给匿名结构体字段赋值
 */
type case_5_1 struct {
	name string
	age int
}
type case_5_2 struct {
	id string
	score float64
}
type case_5_3 struct {
	case_5_1
	case_5_2
}
//和普通结构体创建时赋值一样，注意好嵌套的顺序和层级关系即可
func case_5(){
	c1 := case_5_3{case_5_1{"hello golang",16}, case_5_2{"2021041301",88.5}}
	c2 := case_5_3{
		case_5_1{
			age : 18, //指定字段时就无需按照结构体里字段的顺序进行赋值了
			name : "hello go",
		},
		case_5_2{
			id : "2021041302",
			score : 90.7,
		},
	}
	fmt.Println("c1:",c1,"c2:",c2)
}


/**
嵌入匿名结构体的指针，效率更高
*/
type case_6_1 struct {
	name string
	age int
}
type case_6_2 struct {
	id string
	score float64
}
type case_6_3 struct {
	*case_6_1
	*case_6_2
}
//在创建时传入case_93_1和case_93_2的地址即可
func case_6(){
	c1 := case_6_3{&case_6_1{"hello",20},&case_6_2{"2021041303",60.5}}
	c2 := case_6_3{
		&case_6_1{
			age : 19, //指定字段时就无需按照结构体里字段的顺序进行赋值了
			name : "golang",
		},
		&case_6_2{
			id : "2021041304",
			score : 70.5,
		},
	}
	fmt.Println("c1:",c1,"c2:",c2)
	//c1和c2存储的都是case_6_1和case_6_2两个嵌入匿名结构体的指针，取值要使用指针变量，并且嵌入的匿名结构体指针取值要一个一个取
	fmt.Println("c1:",*c1.case_6_1,*c1.case_6_2,"c2:",*c2.case_6_1,*c2.case_6_2)
}


/**
结构体嵌套的匿名结构体是一个数据类型
那么不仅是结构体数据类型，其他数据类型也可以匿名嵌入结构体，访问时和嵌入的匿名结构体访问方式相同
 */
type case_7_1 struct {
	name string
	age int
}
type case_7_2 struct {
	case_7_1
	int //这里的int和case_7_1是一样的作用，也是个case_7_2里嵌套的匿名数据类型，但int里没有字段，只有一个变量名作为case_7_2的字段，int数据类型，int只能有一个
	number int //如果有第二个int字段就必须给其指定名字，否则两个int会报错
}
func case_7(){
	c1 := case_7_2{}
	c1.name = "hello"
	c1.age = 18
	c1.int = 108
	c1.number = 56
	fmt.Printf("c1:%v\nc1数据类型：%T\nc1.int数据类型：%T",c1,c1,c1.int)
}