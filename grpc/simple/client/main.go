package main

import (
	"context"
	"fmt"

	"github.com/tuanliang/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {
	// 建立网络连接
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// grpc 为我们生成一个客户端调用服务端的sdk
	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "shiyi"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
