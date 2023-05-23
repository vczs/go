package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-case/hello-server/proto"
	"log"
)

type ClientTokenAuth struct{}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "123456",
		"appkey": "golang",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 连接服务
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	conn, err := grpc.Dial("127.0.0.1:8888", opts...)
	if err != nil {
		log.Fatal("failed conn grpc server:", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal("failed conn close:", err)
		}
	}()

	// 创建客户端
	client := service.NewSayHelloClient(conn)

	// 调用服务端api
	res, _ := client.SayHello(context.TODO(), &service.HelloRequest{Name: "张三"})

	fmt.Println(res)
}
