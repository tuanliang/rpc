package main

import (
	"context"
	"fmt"
	"io"
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
func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello, %s", req.Value)}, nil
}

func (s *HelloServiceServer) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		// 接受请求
		req, err := stream.Recv()
		if err != nil {
			// 当前客户端退出
			if err == io.EOF {
				log.Printf("client closed")
				return nil
			}
			return err

		}
		resp := &pb.Response{Value: fmt.Sprintf("hello,%s", req)}
		// 响应请求
		err = stream.Send(resp)
		if err != nil {
			if err == io.EOF {
				log.Printf("client closed")
				return nil
			}
			return err
		}
	}
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
