package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体的每个字段上可以写个tar，该tar可以通过反射机制获取，常见的应用场景有序列化和反序列化
 */
type People struct {
	Name string `json:"name"`  //json:"name"就是给Name字段加了个`json:"name"`的tar，tar必须要写在``内
	Age int `json:"age"` //当结构体字段加了tar时，json解析时就会解析为tar里定义的字段名
	Number string `json:"number"` //tar的格式：`json:"XXX"`
}
func main(){
	people := People{"golang",21,"a1b2c3d4"}
	//将people序列化为json格式
	//json.Marshal()函数里使用了反射
	peopleStr , err := json.Marshal(people)
	if err != nil {
		fmt.Println("出现错误：",err)
	}else{
		fmt.Println(string(peopleStr))
	}
}