package 单元测试入门

func addNumber(num int) int {
	res := 0
	for i := 0 ; i < num ; i++ {
		res += i
	}
	return res
}