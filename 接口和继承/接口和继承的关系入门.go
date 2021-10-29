package main

import "fmt"

/**
接口和继承的关系：从设计层面来说，继承是一种模板设计，接口是一种行为的规范，接口是不破坏继承关系并对继承的一种补充。
继承：解决代码的复用性和可维护性
接口：设计好某种行为规范，让其他对象实现这种行为规范
1.接口比继承更灵活：继承是is-a的关系，接口是like-a的关系。
2.接口在一定程度上实现了代码的解耦。
*/

/**
有一个猴子名叫悟空，不仅继承了猴子会爬树的基因，还学会了鸟儿会飞，鱼儿会游泳的本领
*/
//创建一个猴子
type Monkey struct{
	name string
}
//猴子有爬树的本领
func (monkey Monkey) climbTree(){
	fmt.Println(monkey.name,"会爬树")
}
//创建一个鸟儿
type bird interface{
	fly() //鸟儿有会飞的本领
}
//创建一个鸟儿
type fish interface{
	swimming() //鱼儿有会游泳的本领
}
//创建一个悟空，继承猴子
type wuKong struct{
	Monkey
}
//让悟空学会鸟儿会飞的本领
func (wuKong wuKong) fly() {
	fmt.Println(wuKong.name,"会飞")
}
//让悟空学会鱼儿会游泳的本领
func (wuKong wuKong) swimming() {
	fmt.Println(wuKong.name,"会游泳")
}
//让猴子学会鸟儿和鱼儿的本领
func main(){
	//实例化悟空
	wuKong := wuKong {
		Monkey{
			name : "悟空",
		},
	}
	//悟空继承猴子会爬树的基因
	wuKong.climbTree()
	//让鸟儿教悟空飞
	var bird bird
	bird = wuKong
	bird.fly()
	//让鱼儿教悟空爬树
	var fish fish
	fish = wuKong
	fish.swimming()
}