package main

import "fmt"

/**
多态：变量（实例）具有多种形态。
在GO中多态特征是通过接口实现的：统一的接口调用不同的实现，这时接口变量就呈现不同的形态，即就是多态。
接口体现多态的两种形式：多态参数 多态数组
*/

//创建一个接口
type Usb interface {
	Start() //Usb设备开始工作
	Stop() //Usb设备停止工作
}
//创建Usb手机设备
type Phone struct {
	name string
}
func (phone Phone)Start(){
	fmt.Println("手机开始工作。。。")
}
func (phone Phone)Stop(){
	fmt.Println("手机停止工作。。。")
}
//创建Usb相机设备
type Camera struct {
	name string
}
func (camera Camera)Start(){
	fmt.Println("相机开始工作。。。")
}
func (camera Camera)Stop(){
	fmt.Println("相机停止工作。。。")
}
func main(){
	//创建一个接口类型的数组，里面可以存放Phone和Camera两种实现了Usb接口的结构体变量，体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"PhoneOne"}
	usbArr[1] = Phone{"PhoneTwo"}
	usbArr[2] = Camera{"CameraOne"}
	fmt.Println(usbArr)
}
