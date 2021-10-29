package main

import (
	"encoding/json"
	"fmt"
)

/**
序列化
go自带包json里的Marshal()方法将对象序列化成字符串切片
 */
func main(){
	case_1()
	case_2()
	case_3()
	case_4()
}


/**
结构体序列化
结构体字段首字母必须大写，不然不能序列化
因为json是另外一个包，序列化这里的结构体字段是跨包处理，该结构体字段要被其他包访问就得首字母大写
 */
type people struct {
	Name string `json:"people_name"` //给Name加个json标签，当结构体被序列化时通过反射机制会将Name改为people_name进行序列化，返回的字符串中Name字段为json标签中的内容
	Age int `json:"age"`
	Address string `json:"address"`
}
func case_1(){
	people := people{
		Name:"golang",
		Age:12,
		Address:"Beijing",
	}
	//将p序列化
	jsonStruct , err := json.Marshal(people) //go自带包json里的Marshal()方法将结构体序列化成字符串切片
	if err != nil {
		fmt.Printf("序列化错误，err:=%v",err)
		return
	}
	fmt.Println(string(jsonStruct))  //返回的是切片  要记得转换成字符串类型
}


/**
map序列化
 */
func case_2(){
	var people map[string]interface{}
	people = make(map[string]interface{})
	people["name"] = "golang"
	people["age"] = 18
	people["address"] =  [2]string{"Canada","China"}
	jsonMap , err := json.Marshal(people) //go自带包json里的Marshal()方法将map序列化成字符串切片
	if err != nil {
		fmt.Printf("序列化错误，err:=%v",err)
		return
	}
	fmt.Println(string(jsonMap))  //返回的是切片  要记得转换成字符串类型
}


/**
切片序列化
 */
func case_3(){
	var people []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "hello one"
	m1["age"] = 10
	m1["address"] = "US"
	people = append(people,m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "hello two"
	m2["age"] = 20
	m2["address"] = [2]string{"Canada","China"}
	people = append(people,m2)
	jsonSlice , err := json.Marshal(people) //go自带包json里的Marshal()方法将切片序列化成字符串切片
	if err != nil {
		fmt.Printf("序列化错误，err:=%v",err)
		return
	}
	fmt.Println(string(jsonSlice))  //返回的是切片  要记得转换成字符串类型
}


/**
基本数据类型序列化
没有实际意义
 */
func case_4(){
	dec := 12.51
	jsonDec , err := json.Marshal(dec) //go自带包json里的Marshal()方法将基本数据类型序列化成字符串切片
	if err != nil {
		fmt.Printf("序列化错误，err:=%v",err)
		return
	}
	fmt.Println(string(jsonDec))  //返回的是切片  要记得转换成字符串类型
}