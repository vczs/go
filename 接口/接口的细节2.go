package main

import "fmt"

type AInterface interface {
	sing()
}
type BInterface interface {
	jump()
}
type Kun struct {
	name string
}
func (kun Kun) sing(){
	fmt.Println(kun.name,"在唱！")
}
func (kun Kun) jump(){
	fmt.Println(kun.name,"在跳！")
}
func main(){
	kun := Kun{"坤坤"}
	var a AInterface
	var b BInterface
	a = kun //一个自定义类型可以实现多个接口
	b = kun //一个自定义类型可以实现多个接口
	a.sing()
	b.jump()
}