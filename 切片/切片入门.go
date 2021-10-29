package main

import "fmt"

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
	case_5()
	case_6()
}

/**
切片的初始化
 */
func case_1(){
	//一：定义一个切片直接截取一个已经创建好的数组
	array := [3]int{7,8,9} //创建一个数组
	var slice_a []int = array[0:3]  //创建切片并引用（截取）创建好的数组  也可以写成：slice_a := array[:]    []int可以不写
	//二：make关键字创建切片
	var slice_b []int = make([]int,3,5) //[]int切片数据类型,2是切片长度,4是切片容量,容量必须大于或等于长度    []int可以不写
	slice_b[0] = 7
	slice_b[2] = 8
	slice_b[1] = 9
	//三：定义一个切片,直接指定具体元素
	var slice_c []int = []int{7,8,9}
	fmt.Println("slice_a：",slice_a,len(slice_a),cap(slice_a))  //len()长度,cap()容量
	fmt.Println("slice_b：",slice_b,len(slice_b),cap(slice_b))
	fmt.Println("slice_c：",slice_c,len(slice_c),cap(slice_c))
}


/**
切片截取一个新的切片
 */
func case_2(){
	slice := []int{2,5,1,4}
	newSlice := slice[1:]
	fmt.Println("slice：",slice,len(slice),cap(slice))
	fmt.Println("newSlice：",newSlice,len(newSlice),cap(newSlice))
}

/**
字符串切片
 */
func case_3(){
	var slice []string = []string{"aaa","bbb","ccc","ddd"}
	fmt.Println("slice：",slice,len(slice),cap(slice))
}


/**
切片元素追加
 */
func case_4(){
	slice1 := make([]int,2,4)
	slice1 = []int{1,2,3,4,5}
	slice2 := append(slice1,6,7,8) //也可以用slice1接收,表示在slice1基础上增加6,7,8元素成为新的slice1
	slice3 := append(slice1[0:2],slice1[3:]...)
	fmt.Print("slice2：",slice2,len(slice2),cap(slice2))//追加后容量自己扩容，长度为原切片长度加上追加的元素数量
	fmt.Print("\nslice3：",slice3,len(slice3),cap(slice3))
}


/**
切片复制
 */
func case_5(){
	slice := []string{"aaa","bbb","ccc","ddd"}
	var tarSlice []string = make([]string,3,5)
	copy(tarSlice,slice) //长度不够，只复制能够存储的元素，剩下的省略不复制，不会报错
	fmt.Print("slice：",tarSlice,len(tarSlice),cap(tarSlice))
}


/**
字符串转切片，切片截取字符串，切片转字符串
 */
func case_6(){
	var str = "hello golang!"
	slice := []rune(str[1:])  //字符用[]byte也可以    全部转换：slice := []rune(str)
	slice[5] = 'A'  //更改转换后的切片的元素
	str = string(slice)//再将切片转为字符串打印检查元素是否改变
	fmt.Print("str：",str)
}