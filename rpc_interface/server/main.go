package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/tuanliang/rpc/rpc_interface/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloService)(nil)

// service handler
type HelloService struct {
}

// request收到name，response相应hello，name
func (s *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello,%s", request)
	return nil
}

func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(service.SERVICE_NAME, new(HelloService))

	// 准备socket
	// 然后建立一个唯一的TCP连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listener tcp error:", err)
	}

	// 获取连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go rpc.ServeConn(conn)
	}
}
