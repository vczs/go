package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-case/hello-server/proto"
	"log"
	"net"
)

type server struct {
	service.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	// 校验客户端
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("客户端未传token！")
	}
	var appid, appkey string
	if v, ok := md["appid"]; ok {
		appid = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appkey = v[0]
	}
	if appid != "123456" || appkey != "golang" {
		return nil, errors.New("token不正确！")
	}
	// 业务逻辑
	fmt.Println("client:", req.Name)
	return &service.HelloResponse{Res: "hello," + req.Name}, nil
}

func main() {
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 注册我们自定义的服务到grpc服务
	service.RegisterSayHelloServer(grpcServer, &server{})

	// 监听端口
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("failed listen:", err)
	}

	// 启动grpc服务
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed start grpc server:", err)
	}
}
