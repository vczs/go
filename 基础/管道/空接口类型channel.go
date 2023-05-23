package main

import "fmt"

func main(){
	case_3()
}

/**
如果要在channel内放入不同类型的数据，可以将channel的类型设置为interface{}类型（即空接口类型）
当空接口类型的channel放入结构体类型数据，取出时要类型断言，否则编译不通过。
 */
type People struct{
	Name string
	Age int
}
func case_3(){
	//声明channel并放入数据
	interfaceChan := make(chan interface{},4)
	interfaceChan <- People{"golang",20} //放入一个People类型结构体
	interfaceChan <- 15  //放入一个int类型数据
	interfaceChan <- "你好，GO！" //放入一个string类型数据
	chanMap := make(map[int]string,2) //声明一个map
	chanMap[5] = "hello" //给map赋值
	chanMap[9] = "golang"
	interfaceChan <- chanMap //放入一个map类型数据

	//channel是FIFO的存放原则，所以第一个取出的就是People类型结构体
	//取出People结构体数据
	people := <- interfaceChan //此时people不是结构体数据类型，而是interface{}类型
	fmt.Printf("people:%v\n",people)//直接输出是可以的
	//但是访问people.Name时就会编译不通过，因为此时的people不是结构体类型，而是interface{}类型，编译器找不到people的Name字段
	peopleTem := people.(People) //类型断言people是不是People结构体类型，用peopleTem接收
	fmt.Printf("peopleTem.Name:%v\n",peopleTem.Name) //类型断言people是People结构体类型后就可以访问它的Name字段

	numTem := <- interfaceChan //取出第二个int数据
	strTem := <- interfaceChan //取出第三个string数据
	mapTem := <- interfaceChan //取出第四个map数据
	fmt.Printf("numTem=%v  strTem=%v  mapTem=%v\n",numTem,strTem,mapTem)
	//同样map的具体元素同样也需要类型断言后才可访问
	mapTemTem := mapTem.(map[int]string) //mapTem不是map[int]string类型，而是空接口类型，需要类型断言为map[int]string类型
	fmt.Printf("mapTemTem[5]=%v  mapTemTem[6]=%v\n",mapTemTem[5],mapTemTem[9]) //此时mapTemTem是map[int]string类型，可以访问他的元素
}