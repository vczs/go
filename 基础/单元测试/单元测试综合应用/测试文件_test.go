package 单元测试综合应用

import (
	"testing"
)


func TestMach(t *testing.T){
	people := &People{
		Name : "golang",
		Age : 21,
		Address : "Canada",
	}
	if err := people.mach() ; err != nil {
		t.Fatalf("***错误日志：文件序列化写入失败：%v***",err)
	}
	t.Logf("日志,文件序列化写入成功！！！")
}


func TestUnMach(t *testing.T){
	var people People
	if err := people.unMach() ; err != nil {
		t.Fatalf("***错误日志：文件序列化读取失败：%v***",err)
	}
	t.Logf("日志：文件序列化读取成功！！！")
}