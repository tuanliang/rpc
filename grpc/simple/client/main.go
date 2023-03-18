package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tuanliang/rpc/grpc/middleware/server"
	"github.com/tuanliang/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 建立网络连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// grpc 为我们生成一个客户端调用服务端的sdk
	client := pb.NewHelloServiceClient(conn)

	// 添加认证凭证信息
	crendential := server.NewClientCredential("admin", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), crendential)
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
