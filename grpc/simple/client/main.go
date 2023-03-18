package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tuanliang/rpc/grpc/middleware/client"
	"github.com/tuanliang/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 第二种添加认证凭证信息
	crendital := client.NewAuthentication("admin", "123456")

	// 建立网络连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(crendital))
	if err != nil {
		panic(err)
	}

	// grpc 为我们生成一个客户端调用服务端的sdk
	client := pb.NewHelloServiceClient(conn)

	// 第一种添加认证凭证信息
	// crendential := server.NewClientCredential("admin", "123456")
	// ctx := metadata.NewOutgoingContext(context.Background(), crendential)

	// 第二种添加认证凭证信息1
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs())
	resp, err := client.Hello(ctx, &pb.Request{Value: "shiyi"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// stream, err := client.Channel(context.Background()) //未携带信息，
	stream, err := client.Channel(ctx) //携带信息
	if err != nil {
		panic(err)
	}
	// 启用一个Goroutine发送请求
	go func() {
		for {
			// recover()
			err := stream.Send(&pb.Request{Value: "shiyi1"})
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		// 主循环，负责接收服务端响应
		resp, err = stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}

}
