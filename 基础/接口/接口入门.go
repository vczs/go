package main

import "fmt"

/**
接口：接口是一组仅包含方法名、参数、返回值的未具体实现的方法的集合。
1.一个自定义类型只有实现了某个接口，才能将自身的实例赋值给该接口
2.接口中所有方法都没有方法体，即都是没有实现的方法
3.在go中一个自定义类型需要将某个接口定义的所有方法都实现，才可以说该自定义类型实现了这个接口
4.不仅仅是结构体，只要是自定义数据类型，都可以实现接口
5.GO中的接口不需要显性实现，只要一个变量实现了这个接口中的所有方法，那么就可以说该变量实现了这个接口。
6.一个自定义类型可以实现多个接口
7.go中的接口不能有任何变量，只有未实现的方法
8.一个接口可以继承其他多个接口（继承的接口中不能有相同的方法），如果A接口继承了B和C接口，那么要实现A接口，也必须要将B和C接口的方法全部实现，缺一不可。
9.接口默认是一个指针类型（引用类型），如果没有对接口初始化就使用，就会输出nil
10.一个空的接口没有定义任何方法，可以说所有的类型都实现了该接口。
*/

//创建一个接口
type Usb interface {
	//创建接口的方法（这个接口能做些什么事）
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
	//Working方法负责在有Usb设备接入时通过Usb接口调用Usb设备的Start()和Stop()方法
	//因为Working实现了Usb接口，并且Usb接口拥有Start()和Stop()这两个方法，所以可以通过Usb接口调起Usb设备的Start()和Stop()方法
	usb.Start()
	usb.Stop()
}


/**
创建两个Usb设备
Usb设备有Start()和Stop()两个方法，Usb接口厂家和Usb设备厂家都遵守统一的制作标准，目的就是方便Usb接口和Usb设备的灵活对接
如果一个Usb设备没有Start()和Stop()两个任何一个方法，都不能通过Usb接口接入。（缺少方法的Usb设备在传入该设备对象时编译不通过）
这样一来无论什么Usb设备出厂前都有了Start()和Stop()方法，Usb接口就不用管什么Usb设备，只负责执行该Usb设备Start()和Stop()方法即可
 */
//创建Usb手机设备
type Phone struct {

}
func (phone Phone)Start(){
	fmt.Println("手机开始工作。。。")
}
func (phone Phone)Stop(){
	fmt.Println("手机停止工作。。。")
}
//创建Usb相机设备
type Camera struct {

}
func (camera Camera)Start(){
	fmt.Println("相机开始工作。。。")
}
func (camera Camera)Stop(){
	fmt.Println("相机停止工作。。。")
}


/**
实例化两个Usb设备和Computer并让他们互动
给Usb设备和电脑通电，并将Usb设备通过接口接入电脑
 */
func main(){
	computer := Computer{} //给电脑通电
	phone := Phone{} //给手机通电
	camera := Camera{} //给相机通电
	computer.Working(phone) //手机接入电脑
	computer.Working(camera) //相机接入电脑
}