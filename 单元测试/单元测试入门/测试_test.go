package 单元测试入门

import "testing"

func TestAddNumber(t *testing.T){
	res := addNumber(10)
	if res != 45 {
		t.Fatalf("致命错误：addNumber(10)失败  正确结果=%v  当前结果=%v",45,res)
	}
	t.Logf("日志：addNumber(10)执行正确！！！")
}
