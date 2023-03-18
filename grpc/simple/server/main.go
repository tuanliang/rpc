package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tuanliang/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

// type HelloServiceServer interface {
// 	Hello(context.Context, *Request) (*Response, error)
// 	mustEmbedUnimplementedHelloServiceServer()
// }

type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

// shiyi --> hello,shiyi
func (s *HelloServiceServer) hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello, %s", req.Value)}, nil
}

func main() {
	server := grpc.NewServer()
	// 把实现类注册给Grpc Server
	pb.RegisterHelloServiceServer(server, new(HelloServiceServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	log.Printf("127.0.0.1")
	// 监听socket，HTTP2内置
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
