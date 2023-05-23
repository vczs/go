package main

import "fmt"

//创建一个接口
type Usb interface {
	Start() //Usb设备开始工作
	Stop() //Usb设备停止工作
}

//创建一个可以传入Usb接口的对象
type Computer struct {

}
/**
Computer有一个Working方法拥有Usb接口，其他Usb设备可以通过Working拥有的这个Usb接口接入Computer
*/
func (computer Computer)Working(usb Usb) { //接口默认是引用变量，这里传入的usb对象是地址
	usb.Start()
	if phone , result := usb.(Phone) ; result { //类型断言判断传入的变量参数是不是Phone类型的，是就调用它的Call方法
		phone.Call()
	}
	usb.Stop()
}
//创建Usb手机设备
type Phone struct {
	name string
}
func (phone Phone)Start(){
	fmt.Println("手机开始工作。。。")
}
func (phone Phone)Call(){
	fmt.Println("手机在打电话。。。")
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
	var usbArr [3]Usb
	usbArr[0] = Phone{"phoneOne"}
	usbArr[1] = Camera{"cameraOne"}
	usbArr[2] = Phone{"phoneTwo"}
	var computer Computer
	for _ , v := range usbArr {
		computer.Working(v)
	}
}