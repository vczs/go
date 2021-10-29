package main

import (
	"encoding/json"
	"fmt"
)

/**
反序列化
go自带包json里的Unmarshal()方法将字符串切片反序列化成对应类型
在反序列化一个json字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致（结构体要确保字段完全相同）。
*/
func main(){
	case_5()
	case_6()
	case_7()
}


/**
结构体反序列化
 */
type unPeople struct {
	Name string `json:"people_name"` //给Name加个json标签，当结构体被反序列化时通过反射机制对比字符串内容和结构体字段标签，通过标签将值对应到其结构体字段
	Age int `json:"age"`
	Address string `json:"address"`
}
func case_5(){
	var people unPeople
	str := "{\"people_name\":\"golang\",\"age\":12,\"address\":\"Beijing\"}" //序列化后的字符串
	err  := json.Unmarshal([]byte(str),&people)
	if err != nil {
		fmt.Printf("反序列化错误，err:=%v",err)
		return
	}
	fmt.Println(people)
}


/**
map反序列化
 */
func case_6(){
	var unPeople map[string]interface{}
	//反序列化时map无需make
	//因为Unmarshal()函数里有make()操作，当检测到传入的是map数据类型时就会调用make()操作
	str := "{\"address\":[\"Canada\",\"China\"],\"age\":18,\"name\":\"golang\"}"
	err  := json.Unmarshal([]byte(str),&unPeople)
	if err != nil {
		fmt.Printf("反序列化错误，err:=%v",err)
		return
	}
	fmt.Println(unPeople)
}


/**
切片反序列化
*/
func case_7(){
	var unPeople []map[string]interface{}
	str := "[{\"address\":\"US\",\"age\":10,\"name\":\"hello one\"},{\"address\":[\"Canada\",\"China\"],\"age\":20,\"name\":\"hello two\"}]"
	err  := json.Unmarshal([]byte(str),&unPeople)
	if err != nil {
		fmt.Printf("反序列化错误，err:=%v",err)
		return
	}
	fmt.Println(unPeople)
}

