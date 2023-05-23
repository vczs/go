package 单元测试综合应用

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
}
func (p *People) mach() error {
	//序列化p
	jsonPeople , JsonErr := json.Marshal(p)
	if JsonErr != nil {
		fmt.Printf("jsonPeople序列化失败：%v\n",JsonErr)
		return JsonErr
	}
	fmt.Println("jsonPeople序列化成功！")

	//写入文件
	jsonFilePath := "A:/json.txt"
	if writerErr := ioutil.WriteFile(jsonFilePath,jsonPeople,0) ; writerErr != nil {
		fmt.Printf("文件写入失败：%v\n",writerErr)
		return writerErr
	}

	fmt.Println("文件写入成功")
	return nil
}


func (p *People) unMach() error {
	//读取文件
	jsonFilePath := "A:/json.txt"
	contentSlice , readErr := ioutil.ReadFile(jsonFilePath)
	if readErr != nil {
		fmt.Printf("文件读取发生错误：%v\n",readErr)
		return readErr
	}
	fmt.Printf("文件内容：%v\n",string(contentSlice))

	//反序列化p
	err := json.Unmarshal(contentSlice,p)
	if err != nil {
		fmt.Printf("反序列化错误，err:=%v\n",err)
		return err
	}
	fmt.Printf("反序列化成功：%v\n",*p)
	return nil
}