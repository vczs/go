package hello

import "testing"

func TestSayHello(t *testing.T) {
	res := SayHello("alex")
	t.Logf("SayHello函数结果: %s", res)
}
