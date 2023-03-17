package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/tuanliang/rpc/json_tcp/service"
)

// 约束客户端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	// 建立socket连接
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// 客户端采用json格式来编解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: client,
	}, nil

}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, response *string) error {
	// 执行rpc方法
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var resp string
	if err := c.Hello("shier", &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
