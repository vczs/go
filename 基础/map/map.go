package main

import (
	"fmt"
	"sort"
)

func main(){
	case_1()
	case_2()
	case_3()
	case_4()
	case_5()
	case_6()
	case_7()
	case_8()
}

/**
map的初始化
map：键值对 无序
 */
func case_1(){
	//一：make创建一个map，再给其赋值
	var map_a map[int]string = make(map[int]string) //也可以写成 map_a := make(map[int]string)
	map_a[7] = "aaa"
	map_a[8] = "bbb"
	map_a[9] = "ccc"
	//二:初始化 赋值 一体化
	var map_b map[string]string = map[string]string{"七":"aaa","八":"bbb","九":"ccc"} //也可以写成 map_b := map[int]string{7:"aaa",8:"bbb",9:"ccc"}
	fmt.Print("map_a：",map_a,len(map_a),"\n","map_b：",map_b,len(map_b))
}


/**
map元素查找
 */
func case_2(){
	myMap := make(map[int]string)
	myMap[7] = "aaa"
	myMap[8] = "bbb"
	myMap[9] = "ccc"
	value,result := myMap[5]  //第一个返回值value接收值，第二个返回值result接收查询结果
	if result {
		fmt.Println("result:",result,"value:",value) //查找到result为true，value为查找到的值
	}else{
		fmt.Println("result:",result,"value:",value) //未查找到result为false，value为空
	}
}


/**
map元素删除
*/
func case_3(){
	myMap := make(map[int]string)
	myMap[7] = "aaa"
	myMap[8] = "bbb"
	myMap[9] = "ccc"
	fmt.Println(myMap)
	delete(myMap,8) //第一个参数为map对象，第二个参数要删除的map对象元素的键
	fmt.Println(myMap)
	delete(myMap,1) //没有该键就不做任何操作，也不报错
	fmt.Println(myMap)
}


/**
map元素遍历
go中map的键值都是无规律，所以不能用for循环遍历
map只能用for range语句遍历
注意：go中map是无序的！！！
*/
func case_4(){
	myMap := make(map[int]string)
	myMap[7] = "aaa"
	myMap[8] = "bbb"
	myMap[9] = "ccc"
	for key,value := range myMap{
		fmt.Println("key:",key,"value:",value) //因为map无序，所以遍历打印的顺序不能每次都相同
	}
}


/**
map的嵌套
 */
func case_5(){
	var myMap map[int]map[string]int = make(map[int]map[string]int) //初始化一个map，键是[int]，值是map[string]int
	myMap[5] = make(map[string]int) //每个键的值map也要初始化，初始化后才能添加元素
	myMap[5]["我"] = 22 //添加元素时将myMap[1]看作一个map变量名，["我"]为键，22为值
	myMap[5]["吾"] = 33
	myMap[8] = make(map[string]int) //新的键的值map同样也要初始化才能添加变量
	myMap[8]["他"] = 22
	myMap[8]["她"] = 33
	fmt.Println(myMap) //整个双层map的嵌套打印
	fmt.Println(myMap[5]) //第一层键返回的是第二层的map整体
	fmt.Println(myMap[5]["吾"]) //访问时注意层次即可
}


/**
嵌套的map遍历
map只能用for range语句遍历
 */
func case_6(){
	var myMap map[int]map[string]int = make(map[int]map[string]int)
	myMap[5] = make(map[string]int)
	myMap[5]["我"] = 22
	myMap[5]["吾"] = 33
	myMap[8] = make(map[string]int)
	myMap[8]["他"] = 22
	myMap[8]["她"] = 33
	for key,value := range myMap{ //先遍历myMap的键值对
		fmt.Println("myMap:{ key:",key,"  value:",value,"}")
		for key,value := range value{ //再用获取到的myMap的值遍历myMap嵌套的map的键值对
			fmt.Println("myMap的嵌套:{ key:",key,"  value:",value,"}")
		}
	}
}


/**
map的排序
先遍历map的key并存入一个创建好的slice中，再用sort.Ints(slice)函数对slice的元素进行排序
排序好后的slice的元素即为排序后的map的key，遍历slice的元素并通过map的key获取map的value即完成对map的排序
 */
func case_7(){
	myMap := map[int]string{9:"九",8:"八",5:"五",22:"二十二",12:"十二"}
	var slice []int
	fmt.Println("排序前:")
	for key,value:= range myMap {
		fmt.Println("myMap：{ key:",key,"value:",value,"}")
		slice = append(slice,key) //遍历map并存入slice中
	}
	sort.Ints(slice) //该函数对切片的元素进行排序返回排序好的slice
	fmt.Println("排序后:")
	for _ , value := range slice{ //遍历排序好的slice
		fmt.Println("myMap：{ key:",value,"value:",myMap[value],"}") //slice的value是排序好的map的key，再通过map的key获取map的value，遍历slice就可以得到排好序的map
	}
}


/**
map切片
 */
func case_8(){
	var slice []map[string]string = make([]map[string]string,2,3) //创建一个map[string]string数据类型的切片
	//slice第一个元素
	if slice[0] == nil { //先判断元素是否为nil，为nil就创建map元素
		slice[0] = make(map[string]string) //创建切片的map[string]string元素，先创建才能赋值
	}
	slice[0]["第一个map元素的第一个key"] = "第一个map元素的第一个value" //给第一个slice元素赋值，这里slice[0]相当于map的变量名
	slice[0]["第一个map元素的第二个key"] = "第一个map元素的第二个value"
	//slice第二个元素
	if slice[1] == nil {
		slice[1] = make(map[string]string)
	}
	slice[1]["第二个map元素的第一个key"] = "第二个map元素的第一个value"
	slice[1]["第二个map元素的第二个key"] = "第二个map元素的第二个value"
	fmt.Println("slice:",slice,len(slice),cap(slice))
	for key , value := range slice {
		fmt.Println(key,value)
	}
}